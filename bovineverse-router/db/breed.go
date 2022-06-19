package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"brave-ox-web/model"
)

var (
	BreedDao breedDao = &breedDaoImpl{}
)

type breedDao interface {
	DeleteByTid(tx *gorm.DB, tokenId uint32) error
	QueryCenterList(order uint8, gender uint8, page uint32) (datas []*model.BreedCattleInfo, total uint32, currentPage uint32, err error)
	QueryCenterListByAddr(addr string, order, gender uint8, page uint32) (datas []*model.BreedCattleInfo, total uint32, currentPage uint32, err error)
	QueryBreedInfoByTokenId(id uint32) (*model.BreedCattleInfo, error)
}

type breedDaoImpl struct{}

func (*breedDaoImpl) DeleteByTid(tx *gorm.DB, tokenId uint32) error {
	return tx.Exec("DELETE FROM u_breed_center WHERE token_id = ? ", tokenId).Error
}

func (*breedDaoImpl) QueryCenterList(order uint8, gender uint8, page uint32) ([]*model.BreedCattleInfo, uint32, uint32, error) {
	var condition string
	//u_breed_center AS b
	//u_cattle AS c
	if gender != 0 {
		condition = fmt.Sprint(condition, " c.gender = ", gender)
	}
	return queryBreedCattleData(condition, order, page)
}

func queryBreedCattleData(condition string, order uint8, page uint32) ([]*model.BreedCattleInfo, uint32, uint32, error) {
	if page == 0 {
		page = 1
	}

	q := `
		SELECT b.token_id, b.price, b.seller, c.class, c.is_adult, c.gender, c.gender_seq, c.star,
	    c.life, c.growth, c.energy, c.attack, c.stamina, c.defense, c.milk, c.milk_rate
		FROM u_breed_center AS b 
		INNER JOIN u_cattle AS c ON (b.token_id = c."id")
		WHERE c.dead_at > ?
          `

	countSTMT := `	SELECT count(1) AS total FROM u_breed_center AS b LEFT JOIN u_cattle AS c ON (b.token_id = c."id") WHERE c.dead_at > ?`
	if condition != "" {
		q = fmt.Sprint(q, " AND", condition)
		countSTMT = fmt.Sprint(countSTMT, " AND", condition)
	}
	pageLimit := uint32(50)

	//  一次获取50
	var offset = (page - 1) * pageLimit
	switch order {
	case 0:
		q = fmt.Sprint(q, " ORDER BY b.create_at DESC LIMIT 50 OFFSET ", offset)
	case 1:
		q = fmt.Sprint(q, " ORDER BY b.price ASC LIMIT 50 OFFSET ", offset)
	case 2:
		q = fmt.Sprint(q, " ORDER BY b.price DESC LIMIT 50 OFFSET ", offset)
	default:
		return nil, 0, page, fmt.Errorf("wrong order type:%d", order)
	}

	nows := time.Now().Unix()
	var tmp RawCountType

	result := MainDB().Raw(countSTMT, nows).Scan(&tmp)
	if err := result.Error; err != nil {
		return nil, 0, 0, err
	}

	if tmp.Total <= 0 {
		return nil, 0, page, nil
	}

	page = correctPage(tmp.Total, page, pageLimit)

	rows, err := MainDB().Raw(q, nows).Rows()
	if err != nil {
		return nil, 0, 0, err
	}

	var list []*model.BreedCattleInfo
	for rows.Next() {
		var data model.BreedCattleInfo
		//b.token_id, b.price, b.seller, c.class, c.is_adult, c.gender, c.gender_seq, c.star,
		//	c.life, c.growth, c.energy, c.attack, c.stamina, c.defense, c.milk, c.milk_rate
		rows.Scan(&data.TokenId, &data.Price, &data.Seller, &data.Class, &data.IsAdult, &data.Gender,
			&data.GenderSeq, &data.Star, &data.Life, &data.Growth, &data.Energy, &data.Attack, &data.Stamina,
			&data.Defense, &data.Milk, &data.MilkRate)
		list = append(list, &data)
	}

	return list, tmp.Total, page, nil
}

func (*breedDaoImpl) QueryCenterListByAddr(addr string, order, gender uint8, page uint32) ([]*model.BreedCattleInfo, uint32, uint32, error) {
	//u_breed_center AS b
	//u_cattle AS c
	var condition string
	if gender != 0 {
		condition = fmt.Sprintf(" b.seller = '%s' AND c.gender = %d", addr, gender)
	} else {
		condition = fmt.Sprintf(" b.seller = '%s'", addr)
	}

	return queryBreedCattleData(condition, order, page)
}
func (*breedDaoImpl) QueryBreedInfoByTokenId(id uint32) (*model.BreedCattleInfo, error) {
	//u_breed_center AS b
	//u_cattle AS c
	d, _, _, err := queryBreedCattleData(fmt.Sprintf(" b.token_id = %d", id), 0, 1)
	if err != nil {
		return nil, err
	}

	if len(d) > 0 {
		return d[0], nil
	}
	return nil, nil
}
