package db

import (
	"database/sql"

	"gorm.io/gorm"

	"brave-ox-web/model"
)

var (
	ProfileDao profileDao = &profileDaoImpl{}
)

type profileDao interface {
	ResetPhotoUpdatedNum(address string) error
	QueryWebProfileByAddress(addr string) (*model.ProfileInfo, error)
	TxQueryProfileByAddress(tx *gorm.DB, addr string) (*model.ProfileInfo, error)

	UpdateName(tx *gorm.DB, address string, newName string, updatedNum uint32) bool
	UpdateNameUseFree(tx *gorm.DB, address string, newName string) bool
	UpdateProfilePhoto(address string, photoCode string) error

	TxBatchQueryProfilePhoto(tx *gorm.DB, limit int, offset uint32) ([]*model.ProfileInfo, error)
	TxCheckAndRestAvatar(tx *gorm.DB, tid uint64, expectAddr [2]string) error
	SetGender(address string, gender uint8) (err error, ok bool)
	QueryAddressByName(name string) string
	QueryMulti(addr []string) (data []*model.ProfileInfo, err error)
	QueryRelationship(address string) (data []*model.RelationshipInfo, err error)
}

type profileDaoImpl struct{}

func (*profileDaoImpl) ResetPhotoUpdatedNum(address string) error {
	return MainDB().Table("u_profile").Where("address = ?", address).UpdateColumns(map[string]interface{}{"update_num": 0, "used_free": false}).Error
}

func (this *profileDaoImpl) QueryWebProfileByAddress(addr string) (*model.ProfileInfo, error) {
	tx := MainDB().Begin()
	defer tx.Rollback()
	_, err := this.TxQueryProfileByAddress(tx, addr)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	data := &model.ProfileInfo{}
	stmt := `
			SELECT pro.name,pro.avatar,pro.gender,
 			COALESCE((SELECT p.name FROM u_planet AS p WHERE p.id = m.planet_id),'') AS planet_name,
  			 COALESCE((SELECT p.image FROM u_planet AS p WHERE p.id = m.planet_id),'-1') AS planet_image
			FROM u_profile AS pro LEFT JOIN u_planet_member AS m ON (pro.address = m.player)
			WHERE pro.address = ?
             `
	MainDB().Raw(stmt, addr).Row().Scan(&data.Name, &data.Avatar, &data.Gender, &data.PlanetName, &data.PlanetImage)

	return data, nil
}

func (*profileDaoImpl) TxQueryProfileByAddress(tx *gorm.DB, addr string) (*model.ProfileInfo, error) {
	var data = &model.ProfileInfo{}
	if err := tx.Table("u_profile").Select("address", "name", "avatar", "update_num", "used_free").
		Where("address = ?", addr).First(data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if err = tx.Exec(`INSERT INTO u_profile(address, name, avatar) VALUES(?, ?,'') 
			ON CONFLICT(address) DO NOTHING`, addr, addr).Error; err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (*profileDaoImpl) UpdateName(tx *gorm.DB, address string, newName string, updatedNum uint32) bool {
	result := tx.Exec("UPDATE u_profile SET name = ?, update_num = update_num + 1, update_at = CURRENT_TIMESTAMP WHERE address = ? AND update_num < ?",
		newName, address, updatedNum)
	return result.RowsAffected == 1
}

func (*profileDaoImpl) UpdateNameUseFree(tx *gorm.DB, address string, newName string) bool {
	result := tx.Exec("UPDATE u_profile SET name = ?, used_free = true, update_at = CURRENT_TIMESTAMP WHERE address = ? AND used_free = false",
		newName, address)
	return result.RowsAffected == 1
}

func (*profileDaoImpl) UpdateProfilePhoto(address string, photoCode string) error {
	return MainDB().Exec("UPDATE u_profile SET avatar = ?, update_at = CURRENT_TIMESTAMP WHERE address = ? ", photoCode, address).Error
}

func (*profileDaoImpl) TxBatchQueryProfilePhoto(tx *gorm.DB, limit int, offset uint32) ([]*model.ProfileInfo, error) {
	var list []*model.ProfileInfo
	err := tx.Table("u_profile").Select("address", "avatar").Limit(limit).Offset(int(offset)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}

func (*profileDaoImpl) TxCheckAndRestAvatar(tx *gorm.DB, tid uint64, expectAddr [2]string) error {
	str := `UPDATE u_profile SET avatar = '' WHERE address = ? AND (SELECT count(1) FROM u_cattle WHERE id = ? AND "owner" IN (?,?)) = 0
            AND avatar != ''`
	return tx.Exec(str, expectAddr[0], tid, expectAddr[0], expectAddr[1]).Error
}

func (*profileDaoImpl) SetGender(address string, gender uint8) (err error, ok bool) {
	result := MainDB().Table("u_profile").Where("address = ? AND gender = 0", address).UpdateColumn("gender", gender)
	if result.Error != nil {
		return result.Error, false
	}
	return nil, result.RowsAffected == 1
}

func (*profileDaoImpl) QueryAddressByName(name string) string {
	var addr sql.NullString
	MainDB().Table("u_profile").Select("address").Where("name = ?", name).Row().Scan(&addr)
	return addr.String
}

func (*profileDaoImpl) QueryMulti(addr []string) (data []*model.ProfileInfo, err error) {
	err = MainDB().Table("u_profile").Select([]string{"address", "gender", "name", "avatar"}).Where("address IN ?", addr).Scan(&data).Error
	return
}

func (*profileDaoImpl) QueryRelationship(address string) (data []*model.RelationshipInfo, err error) {
	err = MainDB().Table("u_relationship").Where("? = ANY (pair)", address).Find(&data).Error
	return
}
