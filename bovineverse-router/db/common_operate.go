package db

import (
	"gorm.io/gorm"
)

type RawCountType struct {
	Total uint32 `gorm:"total"`
}

func QueryAll(dest interface{}) error {
	main := MainDB()
	return main.Find(dest).Error
}

func QueryById(m interface{}) error {
	main := MainDB()
	err := main.First(m).Error
	return err
}

func InsertOne(m interface{}) error {
	main := MainDB()
	return main.Create(m).Error
}

func DeleteByModel(m interface{}) error {
	main := MainDB()
	return main.Delete(m).Error
}

func InsertOneWithTx(tx *gorm.DB, m interface{}) error {
	return tx.Save(m).Error
}

func DeleteOneByModelId(tx *gorm.DB, m interface{}) error {
	return tx.Delete(m).Error
}

func CreateBatchByTx(tx *gorm.DB, batch interface{}, size int) error {
	return tx.CreateInBatches(batch, size).Error
}

func correctPage(total uint32, expect uint32, pageLimit uint32) uint32 {
	if expect == 0 {
		return 1
	}
	max := total / pageLimit
	if total%pageLimit > 0 {
		max++
	}
	if max < expect {
		return max
	}
	return expect
}
