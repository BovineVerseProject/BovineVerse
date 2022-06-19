package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"brave-ox-web/log"
	"brave-ox-web/model"
	"brave-ox-web/web/protocol"
)

var (
	AirdropDao airdropDao = &airdropDaoImpl{}
)

type airdropDao interface {
	QueryRewardsByTitleId(titleId uint32) ([]*model.AirdropReward, uint32, error)
	QueryRewardsByAddr(query *protocol.AirdropRequest) ([]*model.AirdropReward, uint32, error)
	QueryTitles() (map[uint32]*model.AirdropTitle, error)
	QueryTitleById(titleId uint32) (*model.AirdropTitle, error)
	QueryRewardById(recordId uint32) (*model.AirdropReward, error)
	UpdateRewardStatus(record *model.AirdropReward, status int) error
	RecoverRewardsStatus() error
	FreezeRewardsByTitleId(ids []uint32) error
	SetFinished(tx *gorm.DB, id uint32) error

	InsertRewards(records *model.AirdropInfo) error
}

type airdropDaoImpl struct{}

func (*airdropDaoImpl) QueryRewardsByTitleId(titleId uint32) ([]*model.AirdropReward, uint32, error) {
	main := MainDB()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	record := make([]*model.AirdropReward, 5)
	err := main.WithContext(ctx).Where("title_id = ?", titleId).Find(&record).Error
	return record, uint32(len(record)), err
}

func (*airdropDaoImpl) QueryRewardsByAddr(query *protocol.AirdropRequest) ([]*model.AirdropReward, uint32, error) {
	main := MainDB()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var total int64

	var stmt string
	var args []interface{}
	if len(query.Address) != 0 {
		stmt += "address = ?"
		args = append(args, query.Address)
	}
	tx := main.WithContext(ctx).Model(&model.AirdropReward{}).Where(stmt, args...).Count(&total)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	list := make([]*model.AirdropReward, 5)
	err := tx.WithContext(ctx).Limit(query.Limit).Offset((query.Page - 1) * query.Limit).Order("id DESC").Find(&list).Error
	return list, uint32(total), err
}

func (*airdropDaoImpl) QueryRewardById(recordId uint32) (*model.AirdropReward, error) {
	record := &model.AirdropReward{}
	err := MainDB().First(record, recordId).Error
	return record, err
}

func (*airdropDaoImpl) UpdateRewardStatus(record *model.AirdropReward, status int) error {
	return MainDB().Model(record).UpdateColumn("status", status).Error
}

func (*airdropDaoImpl) QueryTitleById(titleId uint32) (*model.AirdropTitle, error) {
	record := &model.AirdropTitle{}
	err := MainDB().First(record, titleId).Error
	return record, err
}

func (*airdropDaoImpl) QueryTitles() (map[uint32]*model.AirdropTitle, error) {
	list := make([]*model.AirdropTitle, 5)
	err := MainDB().Model(&model.AirdropTitle{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	records := make(map[uint32]*model.AirdropTitle)
	for _, obj := range list {
		records[obj.Id] = obj
	}
	return records, err
}

func (*airdropDaoImpl) RecoverRewardsStatus() error {
	return MainDB().Model(&model.AirdropReward{}).Where("status = ?", -1).UpdateColumn("status", 0).Error
	// return MainDB().Exec("UPDATE u_airdrop_reward SET status = ?  WHERE status = ? ", 0, -1).Error
}

func (*airdropDaoImpl) FreezeRewardsByTitleId(ids []uint32) error {
	main := MainDB().Begin()
	if main.Error != nil {
		return main.Error
	}
	defer main.Rollback()

	if err := main.Model(&model.AirdropTitle{}).Where("id in ?", ids).UpdateColumn("deadline", true).Error; err != nil {
		return err
	}
	if err := main.Model(&model.AirdropTitle{}).Where("title_id in ? AND status != 1", ids).UpdateColumn("status", 2).Error; err != nil {
		return err
	}

	return main.Commit().Error
}

func (*airdropDaoImpl) SetFinished(tx *gorm.DB, id uint32) error {
	return tx.Table("u_airdrop_reward").Where("id = ?", id).UpdateColumn("status", 1).Error
}

func (a *airdropDaoImpl) InsertRewards(records *model.AirdropInfo) error {
	main := MainDB().Begin()
	if main.Error != nil {
		return main.Error
	}
	defer main.Rollback()

	table := "u_airdrop_reward"

	titleIds := make([]uint32, 0, len(records.Item)+len(records.Usdt)+len(records.Bvg)+len(records.Bvt))
	for _, info := range records.Usdt {
		stmt := fmt.Sprintf("INSERT INTO %s(type, title_id, address, category, from_id, amount) VALUES(%d, ?, ?, %d, %d, ?)", table, 1, 1, 1)
		titleId, err := a.insertAirdrop(main, &info, stmt)
		if err != nil {
			return err
		}
		titleIds = append(titleIds, titleId)
	}
	for _, info := range records.Bvg {
		log.Infof("%+v", info)
		stmt := fmt.Sprintf("INSERT INTO %s(type, title_id, address, category, from_id, amount) VALUES(%d, ?, ?, %d, %d, ?)", table, 2, 1, 2)
		titleId, err := a.insertAirdrop(main, &info, stmt)
		if err != nil {
			return err
		}
		titleIds = append(titleIds, titleId)
	}
	for _, info := range records.Bvt {
		stmt := fmt.Sprintf("INSERT INTO %s(type, title_id, address, category, from_id, amount) VALUES(%d, ?, ?, %d, %d, ?)", table, 3, 1, 3)
		titleId, err := a.insertAirdrop(main, &info, stmt)
		if err != nil {
			return err
		}
		titleIds = append(titleIds, titleId)
	}
	for _, info := range records.Item {
		stmt := fmt.Sprintf("INSERT INTO %s(type, title_id, address, category, from_id, amount) VALUES(%d, ?, ?, %d, %d, ?)", table, 4, 3, 3)
		titleId, err := a.insertAirdrop(main, &info, stmt)
		if err != nil {
			return err
		}
		titleIds = append(titleIds, titleId)
	}

	return main.Commit().Error
}

func (*airdropDaoImpl) insertReward(tx *gorm.DB, stmt string, users []model.AirdropUser, titleId uint32) error {
	for i, user := range users {
		if err := tx.Exec(stmt, titleId, user.Addr, user.Amount).Error; err != nil {
			return err
		}
		log.Infof("[%d] %s，%s, %d", i, stmt, user.Addr, user.Amount)
	}
	return nil
}

func (*airdropDaoImpl) insertTitle(tx *gorm.DB, title *model.AirdropTitle) (uint32, error) {
	begin := time.Now().Unix()
	end := time.Now().AddDate(0, 0, 8).Unix()

	info := model.AirdropTitle{
		TitleCh: title.TitleCh,
		TitleEn: title.TitleEn,
		TitleKr: title.TitleKr,
		DescCh:  title.DescCh,
		DescEn:  title.DescEn,
		DescKr:  title.DescKr,
		Begin:   begin,
		End:     end,
	}
	err := tx.Create(&info).Error
	return info.Id, err
}

func (a *airdropDaoImpl) insertAirdrop(tx *gorm.DB, info *model.Airdrop, stmt string) (uint32, error) {
	titleId, err := a.insertTitle(tx, &info.Title)
	if err != nil {
		return 0, err
	}
	if err := a.insertReward(tx, stmt, info.Users, titleId); err != nil {
		return 0, err
	}
	return titleId, nil
}
