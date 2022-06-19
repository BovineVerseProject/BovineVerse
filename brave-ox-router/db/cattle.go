package db

import (
	"time"

	"gorm.io/gorm"

	"brave-ox-web/model"
)

var CattleDao cattleDao = &cattleDaoImpl{}

type cattleDao interface {
	QueryById(id uint32) (*model.CattleInfo, error)
	QueryList(ids []uint32) ([]*model.CattleInfo, error)
	QueryDeadCattleCount(tx *gorm.DB, address string) (uint32, error)
	UpdateCattleOwner(tx *gorm.DB, id uint32, owner string) error
	SetLifeTime(tx *gorm.DB, id uint32, addTime int64) error
	UpdateToAdult(tx *gorm.DB, id uint32) error
	UpdateCattleStar(tx *gorm.DB, id uint32, star uint8) error
	QueryDeadCattle(ts int64, dest interface{}) error
}

type cattleDaoImpl struct{}

func (*cattleDaoImpl) QueryById(id uint32) (*model.CattleInfo, error) {
	info := &model.CattleInfo{}
	err := MainDB().Table("u_cattle").Where("id = ?", id).First(info).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return info, nil
}

func (*cattleDaoImpl) QueryList(ids []uint32) ([]*model.CattleInfo, error) {
	var list []*model.CattleInfo
	err := MainDB().Table("u_cattle").Where("id IN ?", ids).Find(&list).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return list, nil
}

func (*cattleDaoImpl) QueryDeadCattleCount(tx *gorm.DB, address string) (uint32, error) {
	var count int64
	nows := time.Now().Unix()
	err := tx.Table("u_cattle").Where("owner = ? AND ( dead_at <= ? )", address, nows).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint32(count), err
}

func (*cattleDaoImpl) UpdateCattleOwner(tx *gorm.DB, id uint32, owner string) error {
	return tx.Table("u_cattle").Where("id = ?", id).Update("owner", owner).Error
}

func (*cattleDaoImpl) SetLifeTime(tx *gorm.DB, id uint32, deadTime int64) error {
	return tx.Exec("UPDATE u_cattle SET dead_at = ? WHERE id = ?", id, deadTime).Error
}

func (*cattleDaoImpl) UpdateToAdult(tx *gorm.DB, id uint32) error {
	return tx.Table("u_cattle").Where("id = ?", id).Update("is_adult", true).Error
}

func (*cattleDaoImpl) UpdateCattleStar(tx *gorm.DB, id uint32, star uint8) error {
	return tx.Table("u_cattle").Where("id = ?", id).Update("star", star).Error
}

func (*cattleDaoImpl) QueryDeadCattle(ts int64, dest interface{}) error {
	return MainDB().Where("dead_at <= ?", ts).Find(dest).Error
}

var BurnCattleDao burnCattleDao = &burnCattleDaoImpl{}

type burnCattleDao interface {
	QueryDeadCattleCount(tx *gorm.DB, address string) (uint32, error)
}

type burnCattleDaoImpl struct{}

func (*burnCattleDaoImpl) QueryDeadCattleCount(tx *gorm.DB, address string) (uint32, error) {
	var count int64
	err := tx.Table("r_burn_cattle").Where("owner = ?", address).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint32(count), err
}
