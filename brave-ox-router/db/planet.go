package db

import (
	"fmt"

	"database/sql"

	"gorm.io/gorm"

	"brave-ox-web/model"
)

var (
	PlanetDao planetDao = &planetDaoImpl{}
)

type planetDao interface {
	QueryPlanetsByTypes(types uint32, order uint8, from uint32, prefix string) (datas []*model.PlanetInfo, total uint32, page uint32, err error)
	QueryAllPlanet(order uint8) (datas []*model.PlanetInfo, err error)

	QueryPlanetById(id uint32) (*model.PlanetInfo, error)
	UpdatePlanetName(id uint32, name string, nows int64) (error, bool)
	UpdatePlanetNotice(id uint32, notice string, nows int64) (error, bool)
	QueryMemberByAddress(address string) (member *model.PlanetMember, err error)
	UpdateMemberPos(id uint32, player string, position uint8) (err error, ok bool)
	UpdateImage(id uint32, imageId uint8) error
	QuerySubPlanets(id uint32, from uint32) (datas []*model.PlanetInfo, total uint32, page uint32, err error)
	// QueryMemberPosition(addr string) (pId uint32, pos uint8, joinTime int64, err error)
	QueryMemberPosition(addr string) (members []model.PlanetMember, err error)

	QueryYesterdayMemberGroup(id uint32, less int64) (group []*model.MemberGroup, err error)

	TxSetPlanetMaster(tx *gorm.DB, id uint32, owner string) error
	TxQueryUserPlanetName(tx *gorm.DB, address string) (string, error)
	TxQueryPlanetById(tx *gorm.DB, id uint32) (*model.PlanetInfo, error)
	TxQueryMembers(tx *gorm.DB, id uint32, order uint8, from uint32) (members []*model.PlanetMember, total uint32, currentPage uint32, err error)
	TxSetMasterPosition(tx *gorm.DB, id uint32, player string, pos uint32) error
	TxQueryMemberByAddress(tx *gorm.DB, owner string) (member *model.PlanetMember, err error)
	TxUpdateScore(tx *gorm.DB, player string, num string) error
	TxUpdateFee(tx *gorm.DB, id uint32, fee string) error
	TxUpdatePlanetMaster(tx *gorm.DB, id uint32, master string) error
	TxUpdateMasterPos(tx *gorm.DB, tokenId uint32, master string, pos uint8) error
	TxUpdatePlanetScore(tx *gorm.DB, id uint32, value string) error
	TxUpdatePlanetPopulation(tx *gorm.DB, id uint32) error
	TxFixedPlanetMemberMaster(tx *gorm.DB, id uint32, owner string) error
}

type planetDaoImpl struct{}

func (*planetDaoImpl) QueryPlanetsByTypes(types uint32, order uint8, from uint32, prefix string) (datas []*model.PlanetInfo, total uint32, page uint32, err error) {
	tx := MainDB().Begin()
	defer tx.Rollback()

	countStr := "SELECT COUNT(1) AS total FROM u_planet AS p WHERE p.planet_type = ? "
	if prefix != "" {
		countStr = fmt.Sprint(countStr, " AND p.name LIKE '", prefix, "%'")
	}

	var tmp RawCountType
	tx.Raw(countStr, types).Scan(&tmp)
	if tmp.Total <= 0 {
		return nil, 0, 0, err
	}

	limit := uint32(50)
	page = correctPage(tmp.Total, from, limit)

	stmt := `
			SELECT p."id", p."name", p."notice", p."fee", p."image", u."name" AS masterName, p."master",
			p."population", p."total_score"
            FROM u_planet AS p LEFT JOIN u_profile AS u ON (p."master" = u."address") 
            WHERE p."planet_type" = ?
        	`

	if prefix != "" {
		stmt = fmt.Sprint(stmt, " AND p.name LIKE '", prefix, "%'")
	}

	offset := (page - 1) * limit
	switch order {
	case 0:
		stmt = fmt.Sprint(stmt, " ORDER BY p.id LIMIT 50 OFFSET ", offset)
	case 1:
		stmt = fmt.Sprint(stmt, " ORDER BY population DESC, p.id LIMIT 50 OFFSET ", offset)
	case 2:
		stmt = fmt.Sprint(stmt, " ORDER BY total_score DESC, p.id LIMIT 50 OFFSET ", offset)
	}

	rows, err := tx.Raw(stmt, types).Rows()
	if err != nil {
		return nil, 0, 0, err
	}

	for rows.Next() {
		var data model.PlanetInfo
		var ownerName sql.NullString
		// id,name,notice,fee,image,ownerName,address,population,total_score
		rows.Scan(&data.Id, &data.Name, &data.Notice, &data.Fee, &data.Image, &ownerName, &data.Owner, &data.Population, &data.TotalScore)
		if data.TotalScore == "" {
			data.TotalScore = "0"
		}
		data.OwnerName = ownerName.String
		datas = append(datas, &data)
	}

	return datas, tmp.Total, page, nil
}

func (*planetDaoImpl) QueryAllPlanet(order uint8) (datas []*model.PlanetInfo, err error) {
	tx := MainDB().Begin()
	defer tx.Rollback()

	stmt := `
			SELECT p."id", p."name", p."notice", p."fee", p."image", u."name" AS masterName, p."master",
			p."population", p."total_score"
            FROM u_planet AS p LEFT JOIN u_profile AS u ON (p."master" = u."address")
        	`

	switch order {
	case 0:
		stmt = fmt.Sprint(stmt, " ORDER BY p.id ")
	case 1:
		stmt = fmt.Sprint(stmt, " ORDER BY p.population DESC, p.id")
	case 2:
		stmt = fmt.Sprint(stmt, " ORDER BY p.total_score DESC, p.id")
	}

	rows, err := tx.Raw(stmt).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data model.PlanetInfo
		var ownerName sql.NullString
		// id,name,notice,fee,image,ownerName,address,population,total_score
		rows.Scan(&data.Id, &data.Name, &data.Notice, &data.Fee, &data.Image, &ownerName, &data.Owner, &data.Population, &data.TotalScore)
		data.OwnerName = ownerName.String
		datas = append(datas, &data)
	}

	return datas, nil
}

func (this *planetDaoImpl) QueryPlanetById(id uint32) (*model.PlanetInfo, error) {
	return this.TxQueryPlanetById(MainDB(), id)
}

func (*planetDaoImpl) TxSetPlanetMaster(tx *gorm.DB, id uint32, owner string) error {
	return tx.Table("u_planet").Where("id = ? AND master != ?", id, owner).UpdateColumn("master", owner).Error

	// p := &model.PlanetMember{}
	// if err := tx.Where("planet_id = ? AND player != ?", id, owner).Attrs(model.PlanetMember{PlanetId: id, Player: owner, Position: 4}).FirstOrCreate(p).Error; err != nil {
	// 	log.Error(err)
	// 	return err
	// }
	// return nil
}

func (*planetDaoImpl) TxFixedPlanetMemberMaster(tx *gorm.DB, id uint32, owner string) error {
	p := &model.PlanetMember{}
	return tx.Where("planet_id = ? AND player != ?", id, owner).Attrs(model.PlanetMember{PlanetId: id, Player: owner, Position: 4}).FirstOrCreate(p).Error
}

func (*planetDaoImpl) UpdatePlanetName(id uint32, name string, nows int64) (error, bool) {
	result := MainDB().Table("u_planet").Where("id = ? AND update_name_at < ?", id, nows).
		UpdateColumns(map[string]interface{}{"name": name, "update_name_at": nows})
	if result.Error != nil {
		return result.Error, false
	}
	return nil, result.RowsAffected == 1
}

func (*planetDaoImpl) UpdatePlanetNotice(id uint32, notice string, nows int64) (error, bool) {
	result := MainDB().Table("u_planet").Where("id = ? AND update_notice_at < ?", id, nows).
		UpdateColumns(map[string]interface{}{"notice": notice, "update_notice_at": nows})
	if result.Error != nil {
		return result.Error, false
	}
	return nil, result.RowsAffected == 1
}

func (*planetDaoImpl) TxQueryUserPlanetName(tx *gorm.DB, address string) (string, error) {
	// stmt := "SELECT p.name FROM u_planet AS p WHERE p.id = (SELECT m.planet_id FROM u_planet_member AS m WHERE m.player = ?)"
	stmt := "SELECT p.name FROM u_planet AS p WHERE p.id = (SELECT m.planet_id FROM u_planet_member AS m WHERE m.player = ? AND m.planet_id != (SELECT p.id FROM u_planet AS p where p.planet_type = 2 and p.master = ?))"
	row := tx.Raw(stmt, address, address).Row()
	var n sql.NullString
	err := row.Scan(&n)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return n.String, nil
}

func (this *planetDaoImpl) TxQueryPlanetById(tx *gorm.DB, id uint32) (*model.PlanetInfo, error) {
	var data model.PlanetInfo
	stmt := `
			SELECT p."id", p."name", p."master", p."notice", p."fee", p."update_name_at", p."update_notice_at", 
			p."image", u."name" AS master_name,
			p."population" 
            FROM u_planet AS p LEFT JOIN u_profile AS u ON (p."master" = u."address") 
            WHERE p."id" = ?
        	`

	var ownerName sql.NullString
	tx.Raw(stmt, id).Row().Scan(&data.Id, &data.Name, &data.Owner, &data.Notice, &data.Fee, &data.UpdateNameAt,
		&data.UpdateNoticeAt, &data.Image, &ownerName, &data.Population)
	if data.Id == 0 {
		return nil, nil
	}
	data.OwnerName = ownerName.String
	return &data, nil
}

func (*planetDaoImpl) TxQueryMembers(tx *gorm.DB, id uint32, order uint8, from uint32) ([]*model.PlanetMember, uint32, uint32, error) {
	var total int64
	tx.Table("u_planet_member").Where("planet_id = ?", id).Count(&total)
	if total <= 0 {
		return nil, 0, 0, nil
	}

	page := correctPage(uint32(total), from, 50)
	stmt := `
			SELECT m.planet_id, m.player, m.score, m.position, p.name, p.avatar,
			(SELECT COUNT ( 1 ) AS cattle_num FROM u_cattle AS c 
			WHERE c."owner" = m.player),
			(SELECT COUNT ( 1 ) AS burn_num FROM r_burn_cattle AS b 
			WHERE b."owner" = m.player) 
			FROM u_planet_member AS m
			LEFT JOIN u_profile AS p ON (m.player = p.address)
			WHERE m.planet_id = ?
			`

	offset := (page - 1) * 50
	switch order {
	case 0:
		stmt = fmt.Sprint(stmt, " ORDER BY m.position DESC LIMIT 50 OFFSET ", offset)
	case 1:
		stmt = fmt.Sprint(stmt, " ORDER BY m.score DESC LIMIT 50 OFFSET ", offset)
	}
	rows, err := tx.Raw(stmt, id).Rows()
	if err != nil {
		return nil, 0, 1, err
	}

	var members []*model.PlanetMember
	for rows.Next() {
		var member model.PlanetMember
		var burnCattle uint32
		var nullName sql.NullString
		var nullAvatar sql.NullString
		err = rows.Scan(&member.PlanetId, &member.Player, &member.Score, &member.Position, &nullName, &nullAvatar, &member.Count, &burnCattle)
		if err != nil {
			return nil, 0, 1, err
		}
		member.Count += burnCattle // u_cattle表加burn表
		member.Name = nullName.String
		member.Avatar = nullAvatar.String
		members = append(members, &member)
	}
	return members, uint32(total), page, nil
}
func (*planetDaoImpl) TxSetMasterPosition(tx *gorm.DB, id uint32, player string, pos uint32) error {
	return tx.Exec(`
				UPDATE u_planet_member SET position = ? 
				WHERE player = ? AND  (SELECT COUNT(1) FROM u_planet WHERE id = ? AND master = ?) = 1
                  `, pos, player, id, player).Error
}

func (this *planetDaoImpl) QueryMemberByAddress(addr string) (member *model.PlanetMember, err error) {
	return this.TxQueryMemberByAddress(MainDB(), addr)
}

func (this *planetDaoImpl) QueryMemberPosition(addr string) (members []model.PlanetMember, err error) {
	m := &model.PlanetMember{}

	db := MainDB()
	err = db.Model(m).Select("planet_id,position,create_at").Where("player = ?", addr).Scan(&members).Error
	if err != nil {
		return
	}

	return
}

func (*planetDaoImpl) TxQueryMemberByAddress(tx *gorm.DB, addr string) (member *model.PlanetMember, err error) {
	stmt := `
			SELECT m.planet_id, m.score, m.position, p.name, p.avatar,
			(SELECT COUNT ( 1 ) AS cattle_num FROM u_cattle AS c 
			WHERE c."owner" = m.player),
			(SELECT COUNT ( 1 ) AS burn_num FROM r_burn_cattle AS b 
			WHERE b."owner" = m.player) 
			FROM u_planet_member AS m
			LEFT JOIN u_profile AS p ON (m.player = p.address)
			WHERE m.player = ?
			`
	rows, err := tx.Raw(stmt, addr).Rows()
	if err != nil {
		return nil, err
	}
	var nullName sql.NullString
	var nullAvatar sql.NullString

	member = new(model.PlanetMember)
	var deadCount uint32
	err = rows.Scan(&member.PlanetId, &member.Score, &member.Position, &nullName, &nullAvatar, &member.Count, &deadCount)
	if err != nil {
		return nil, err
	}
	member.Count += deadCount // u_cattle表加burn表
	member.Name = nullName.String
	member.Avatar = nullAvatar.String
	return
}

func (*planetDaoImpl) UpdateMemberPos(id uint32, player string, position uint8) (error, bool) {
	result := MainDB().Table("u_planet_member").Where("planet_id = ? AND player = ? AND position != ?", id, player, position).
		UpdateColumn("position", position)
	if result.Error != nil {
		return result.Error, false
	}
	return nil, result.RowsAffected == 1
}

func (*planetDaoImpl) UpdateImage(id uint32, imageId uint8) error {
	return MainDB().Table("u_planet").Where("id = ? AND image != ?", id, imageId).UpdateColumn("image", imageId).Error
}

func (*planetDaoImpl) QuerySubPlanets(id uint32, from uint32) (datas []*model.PlanetInfo, total uint32, page uint32, err error) {
	tx := MainDB().Begin()
	defer tx.Rollback()
	var count int64
	tx.Table("u_planet").Where("parent_id = ?", id).Count(&count)

	if count <= 0 {
		return
	}

	page = correctPage(uint32(count), from, 8)

	stmt := `
			SELECT p."id", p."name", p."image",
			p."total_score" 
            FROM u_planet AS p
            WHERE p."parent_id" = ? LIMIT 8 OFFSET ?
        	`

	offset := (page - 1) * 8

	rows, err := tx.Raw(stmt, id, offset).Rows()
	if err != nil {
		return nil, 0, 0, err
	}

	for rows.Next() {
		var data model.PlanetInfo
		// id,name,image,total_score
		rows.Scan(&data.Id, &data.Name, &data.Image, &data.TotalScore)
		datas = append(datas, &data)
	}

	return datas, uint32(count), page, nil
}

func (*planetDaoImpl) TxUpdateScore(tx *gorm.DB, player string, num string) error {
	return tx.Exec("UPDATE u_planet_member SET score = score + ? WHERE player = ?", num, player).Error
}

func (*planetDaoImpl) TxUpdateFee(tx *gorm.DB, id uint32, fee string) error {
	return tx.Exec("UPDATE u_planet SET fee = ? WHERE id = ?", fee, id).Error
}

func (*planetDaoImpl) TxUpdatePlanetMaster(tx *gorm.DB, id uint32, master string) error {
	return tx.Exec("UPDATE u_planet SET master = ? WHERE id = ?", master, id).Error
}

func (*planetDaoImpl) TxUpdateMasterPos(tx *gorm.DB, id uint32, master string, pos uint8) error {
	return tx.Table("u_planet_member").Where("planet_id = ? AND player = ?", id, master).Update("position", pos).Error
}

func (*planetDaoImpl) TxUpdatePlanetScore(tx *gorm.DB, id uint32, value string) error {
	return tx.Exec("UPDATE u_planet SET total_score = total_score + ? WHERE id = ?", value, id).Error
}
func (*planetDaoImpl) TxUpdatePlanetPopulation(tx *gorm.DB, id uint32) error {
	return tx.Exec("UPDATE u_planet SET population = population + 1 WHERE id = ?", id).Error
}

func (*planetDaoImpl) QueryYesterdayMemberGroup(id uint32, lessThanTime int64) (group []*model.MemberGroup, err error) {
	err = MainDB().Table("u_planet_member").Select("position, COUNT(*) AS people").
		Where("planet_id = ? AND create_at < ?", id, lessThanTime).Group("position").Scan(&group).Error
	return
}
