package db

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"

	"brave-ox-web/model"
	"brave-ox-web/web/protocol"
)

var (
	MarketDao marketDao = &marketplaceDaoImpl{}
)

const pageLimit = 20

type marketDao interface {
	TxUpdateUPrice(tx *gorm.DB, id uint32, price float64) error

	QuerySellingCattle(cond *protocol.MarketCattleRequest) (datas []*model.MarketCattleInfo, total uint32, err error)
	QueryMarketplaceMysteryBox(query *protocol.MarketRequest) (datas []*model.MarketSaleInfo, total uint32, err error)
	QueryPlanet(order uint8, fromPage uint32) (datas []*model.MarketPlanetInfo, total uint32, page uint32, err error)

	InsertSaleRecord(tx *gorm.DB, record *model.MarketSaleInfo) error
	DeleteByGoodsId(tx *gorm.DB, id uint32) error

	QueryHistory(address string, page uint32) (datas []*model.SaleGoodsRecord, total uint32, currentPage uint32, err error)
	UpdateUPrice(types uint8, rate float64) error
}

type marketplaceDaoImpl struct{}

func (m *marketplaceDaoImpl) TxUpdateUPrice(tx *gorm.DB, id uint32, price float64) error {
	return tx.Table("u_market").Where("id = ? AND u_price != ?", id, price).UpdateColumn("u_price", price).Error
}

func (m *marketplaceDaoImpl) QuerySellingCattle(query *protocol.MarketCattleRequest) (list []*model.MarketCattleInfo, count uint32, err error) {
	stmt := `
				SELECT m.token_id, m.trade_type, m.price, m.u_price, m.seller, c.class, c.gender, c.gender_seq, c.is_adult
				FROM u_market AS m, ( SELECT id AS cid, class, gender, gender_seq, is_adult, star FROM u_cattle
			`

	main := MainDB()
	var topCondition string

	// 条件前面记得加空格!

	// ------------ 类型组--------------------------------------------
	if query.Class != 0 {
		topCondition = fmt.Sprintf(" class = %d AND", query.Class)
	}

	if query.Gender != 0 {
		topCondition = fmt.Sprintf(" %s gender = %d AND", topCondition, query.Gender)
	}

	var l int

	if topCondition != "" {
		topCondition = topCondition[:len(topCondition)-3]
		topCondition = fmt.Sprintf(" (%s)", topCondition)
	}

	var commonStmt string
	if l = len(query.Life); l > 0 {
		life := m.buildRangeSTMT(query.Life, "life")
		commonStmt += life
	}

	if l = len(query.Growth); l > 0 {
		growth := m.buildRangeSTMT(query.Growth, "growth")
		commonStmt += growth
	}

	if l = len(query.Energy); l > 0 {
		hunger := m.buildRangeSTMT(query.Energy, "energy")
		commonStmt += hunger
	}

	if commonStmt != "" {
		commonStmt = commonStmt[:len(commonStmt)-3]
	}

	var oxAttrStmt string
	if l = len(query.Attack); l > 0 {
		attack := m.buildRangeSTMT(query.Attack, "attack")
		oxAttrStmt += attack
	}

	if l = len(query.Stamina); l > 0 {
		stamina := m.buildRangeSTMT(query.Stamina, "stamina")
		oxAttrStmt += stamina
	}

	if l = len(query.Defence); l > 0 {
		defence := m.buildRangeSTMT(query.Defence, "defense")
		oxAttrStmt += defence
	}

	if oxAttrStmt != "" {
		oxAttrStmt = oxAttrStmt[:len(oxAttrStmt)-3]
	}

	var cowAttrStmt string
	if l = len(query.Milk); l > 0 {
		productNum := m.buildRangeSTMT(query.Milk, "milk")
		cowAttrStmt += productNum
	}

	if l = len(query.MilkRate); l > 0 {
		productRate := m.buildRangeSTMT(query.MilkRate, "milk_rate")
		cowAttrStmt += productRate
	}

	if cowAttrStmt != "" {
		cowAttrStmt = cowAttrStmt[:len(cowAttrStmt)-3]
	}

	var genderGroup string
	if oxAttrStmt != "" && cowAttrStmt != "" {
		genderGroup = fmt.Sprintf(" (%sOR%s)", oxAttrStmt, cowAttrStmt)
	} else if oxAttrStmt != "" {
		genderGroup = fmt.Sprintf(" (%s)", oxAttrStmt)
	} else if cowAttrStmt != "" {
		genderGroup = fmt.Sprintf(" (%s)", cowAttrStmt)
	}

	if genderGroup != "" {
		if commonStmt != "" {
			commonStmt = fmt.Sprint(commonStmt, " AND ", genderGroup)
		} else {
			commonStmt = genderGroup
		}
	}

	if topCondition != "" {
		if commonStmt != "" {
			topCondition = fmt.Sprint(topCondition, " AND", commonStmt)
		}
	} else {
		topCondition = commonStmt
	}

	now := time.Now().Unix()
	if topCondition != "" {
		topCondition = fmt.Sprintf("%s AND dead_at > %d", topCondition, now)
	} else {
		topCondition = fmt.Sprintf("%s dead_at > %d", topCondition, now)
	}
	topCondition = fmt.Sprint(" WHERE", topCondition, ") AS c")

	stmt = fmt.Sprint(stmt, topCondition)
	countStmt := fmt.Sprint("SELECT COUNT(1) AS total FROM u_market AS m, ( SELECT id AS cid FROM u_cattle", topCondition)

	marketCond := " WHERE m.goods_type = 1 AND m.token_id = c.cid"
	stmt = fmt.Sprint(stmt, marketCond)
	countStmt = fmt.Sprint(countStmt, marketCond)

	var tmp RawCountType

	result := main.Raw(countStmt).Scan(&tmp)
	if err = result.Error; err != nil {
		return
	}
	if tmp.Total <= 0 {
		return
	}

	sortMethod := " ORDER BY u_price"
	switch query.Sort {
	case 1:
		sortMethod += " DESC"
	case 2:
		sortMethod = " ORDER BY block_time DESC"
	}

	stmt = fmt.Sprint(stmt, sortMethod, " ,m.token_id")

	count = tmp.Total
	query.Page = correctPage(count, query.Page, pageLimit)
	var offset = (query.Page - 1) * pageLimit

	stmt += fmt.Sprintf(" LIMIT %d OFFSET %d", pageLimit, offset)

	rows, err := main.Raw(stmt).Rows()
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var saleInfo model.MarketCattleInfo
		var seller sql.NullString
		err = rows.Scan(&saleInfo.TokenId, &saleInfo.TradeType, &saleInfo.Price, &saleInfo.UPrice, &seller, &saleInfo.Class, &saleInfo.Gender, &saleInfo.GenderSeq, &saleInfo.IsAdult)
		if err != nil {
			return nil, 0, err
		}
		saleInfo.Seller = seller.String
		list = append(list, &saleInfo)
	}

	return
}

func (*marketplaceDaoImpl) buildMultiSTMT(array []uint32, length int, field string) string {
	var stmt string
	if length == 1 {
		stmt = fmt.Sprintf(" %s = %d OR", field, array[0])
	} else {
		var star string
		for _, v := range array {
			star = fmt.Sprintf("%s%d,", star, v)
		}
		star = star[:len(star)-1]
		stmt = fmt.Sprintf(" %s IN (%s) OR", field, star)
	}

	return stmt
}

func (*marketplaceDaoImpl) buildRangeSTMT(array []uint32, field string) string {
	if array[0] == array[1] {
		return fmt.Sprintf(" %s = %d AND", field, array[0])
	}
	if array[1] == 15000 {
		return fmt.Sprintf(" %s >= %d AND", field, array[0])
	}

	return fmt.Sprintf(" (%s >= %d AND %s <= %d) AND", field, array[0], field, array[1])
}

func (*marketplaceDaoImpl) QueryMarketplaceMysteryBox(query *protocol.MarketRequest) ([]*model.MarketSaleInfo, uint32, error) {
	main := MainDB()

	var total int64
	info := &model.MarketSaleInfo{}
	whereCond := "goods_type = 2"
	err := main.Model(info).Where(whereCond).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if total <= 0 {
		return nil, 0, nil
	}

	sortMethod := " u_price"
	switch query.Sort {
	case 1:
		sortMethod += " DESC"
	case 2:
		sortMethod = " block_time DESC"
	}

	sortMethod = fmt.Sprint(sortMethod, " ,token_id")

	page := correctPage(uint32(total), query.Page, pageLimit)
	query.Page = page
	offset := (page - 1) * pageLimit
	var dest []*model.MarketSaleInfo
	err = main.Model(info).Select("token_id", "trade_type", "price", "u_price", "seller").
		Where(whereCond).Order(sortMethod).Limit(pageLimit).Offset(int(offset)).Find(&dest).Error
	if err != nil {
		return nil, 0, err
	}

	return dest, uint32(total), nil
}

func (*marketplaceDaoImpl) InsertSaleRecord(tx *gorm.DB, data *model.MarketSaleInfo) error {
	stmt := `INSERT INTO u_market(id,goods_type,token_id,trade_type,price,block_time,seller)
		 VALUES(?,?,?,?,?,?,?) ON CONFLICT (id) DO NOTHING`
	return tx.Exec(stmt, data.Id, data.GoodsType, data.TokenId, data.TradeType, data.Price, data.BlockTime, data.Seller).Error
}

func (*marketplaceDaoImpl) DeleteByGoodsId(tx *gorm.DB, id uint32) error {
	return tx.Exec("DELETE FROM u_market WHERE id = ?", id).Error
}

func (*marketplaceDaoImpl) QueryPlanet(order uint8, fromPage uint32) ([]*model.MarketPlanetInfo, uint32, uint32, error) {
	main := MainDB()

	var total int64
	info := &model.MarketSaleInfo{}
	whereCond := "goods_type = 3"
	err := main.Model(info).Where(whereCond).Count(&total).Error
	if err != nil {
		return nil, 0, 0, err
	}

	if total <= 0 {
		return nil, 0, 0, nil
	}
	stmt := `
			SELECT m.token_id, p.planet_type, m.trade_type, m.price, m.u_price, m.seller FROM u_market AS m 
			INNER JOIN u_planet721 AS p ON (m.token_id = p.id)
			WHERE m.goods_type = 3
            `

	page := correctPage(uint32(total), fromPage, pageLimit)
	offset := (page - 1) * pageLimit
	var sortMethod string
	switch order {
	case 0:
		sortMethod = " ORDER BY m.u_price ASC"
	case 1:
		sortMethod = " ORDER BY m.u_price DESC"
	case 2:
		sortMethod = " ORDER BY m.block_time DESC"
	}

	stmt = fmt.Sprint(stmt, sortMethod, " , m.token_id", " LIMIT ", pageLimit, " OFFSET ?")

	rows, err := main.Raw(stmt, offset).Rows()
	if err != nil {
		return nil, 0, 0, nil
	}

	var datas []*model.MarketPlanetInfo
	for rows.Next() {
		var data model.MarketPlanetInfo
		var seller sql.NullString
		rows.Scan(&data.TokenId, &data.Types, &data.TradeType, &data.Price, &data.UPrice, &seller)
		data.Seller = seller.String
		datas = append(datas, &data)
	}

	return datas, uint32(total), page, nil
}

func (*marketplaceDaoImpl) QueryHistory(address string, page uint32) (datas []*model.SaleGoodsRecord, total uint32, currentPage uint32, err error) {
	countSQL := " SELECT COUNT(1) AS total FROM r_market WHERE seller = ? OR buyer = ?"
	var tmp RawCountType

	if page == 0 {
		page = 1
	}

	result := MainDB().Raw(countSQL, address, address).Scan(&tmp)
	if err = result.Error; err != nil {
		return
	}
	total = tmp.Total

	pageLimit := uint32(10)
	currentPage = correctPage(tmp.Total, page, pageLimit)
	if currentPage <= 0 {
		return
	}
	err = MainDB().Table("r_market").Where("seller = ? OR buyer = ?", address, address).
		Order("block_time DESC").Limit(int(pageLimit)).Offset(int((page - 1) * pageLimit)).Find(&datas).Error
	return
}

func (*marketplaceDaoImpl) UpdateUPrice(types uint8, rate float64) error {
	stmt := `
			UPDATE u_market SET u_price = CEIL(price * ? * 1000 :: numeric) / 1000 WHERE trade_type = ?
			`
	return MainDB().Exec(stmt, rate, types).Error
}
