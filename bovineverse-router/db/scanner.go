package db

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"brave-ox-web/log"
)

var (
	ScannerDao scannerDao = &scannerDaoImpl{}
)

type scannerDao interface {
	FetchIncompleteBlocks(chain string, limit int, dest interface{})
	UpdateBlockDispatchTs(chain string, block int64, ts int64)
	UpdateBlockStatus(chain string, blockTime int64, block uint64) (err error)
	IncreaseFailTimes(chain string, block int64)
	GetKnownLatest(chain string, dest interface{})
	BatchInsertBlock(chain string, blocks []uint64)
	GetKnownOldest(chain string, dest interface{})
	GetKnownOldestDone(chain string, dest interface{})
	CleanOldBlock(chain string, threshold uint64) bool
	ExistTxHash(chain, hash string) bool
	NewBlockHash(tx *gorm.DB, chain string, txHash string, t time.Time) error
	Clear30DaysTx() (uint32, error)
}

type scannerDaoImpl struct{}

func (*scannerDaoImpl) FetchIncompleteBlocks(chain string, limit int, dest interface{}) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	//now := time.Now().Unix()

	//stmt := fmt.Sprintf("SELECT `block`, `fail_times` FROM %s WHERE `done`=? AND `dispatch_ts`<? AND `fail_times` < 5 ORDER BY `fail_times` ASC, `block` DESC LIMIT ?", blockTable)
	db.Table(blockTable).Select("block").Where("done = false  AND fail_times < 5 ").
		Order(fmt.Sprintf("block LIMIT %d", limit)).Find(dest)
}

func (*scannerDaoImpl) UpdateBlockDispatchTs(chain string, block int64, ts int64) {
	db := MainDB()
	err := db.Table(fmt.Sprintf("s_%s_block", chain)).Where("block = ?", block).Update("dispatch_ts", ts).Error
	if err != nil {
		log.Errorf("scanner.UpdateBlockDispatchTs err:%v", err)
	}
}

func (*scannerDaoImpl) UpdateBlockStatus(chain string, blockTime int64, block uint64) (err error) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)

	return db.Table(blockTable).Where("block = ? AND done = false", block).
		Updates(map[string]interface{}{"done": true, "time": blockTime}).Error
}

func (*scannerDaoImpl) IncreaseFailTimes(chain string, block int64) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	stmt := fmt.Sprintf("UPDATE %s SET fail_times=fail_times + 1 WHERE block=? AND fail_times < 5", blockTable)
	db.Exec(stmt, block)
}

func (*scannerDaoImpl) GetKnownLatest(chain string, dest interface{}) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	db.Table(blockTable).Select(" MAX(block)").Row().Scan(dest)
}

func (*scannerDaoImpl) BatchInsertBlock(chain string, blocks []uint64) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	var vals []interface{}
	stmt := fmt.Sprintf("INSERT INTO %s(block) VALUES", blockTable)
	for _, block := range blocks {
		stmt += "(?),"
		vals = append(vals, block)
	}
	stmt = fmt.Sprint(strings.TrimSuffix(stmt, ","), " ON CONFLICT(block) DO NOTHING")
	err := db.Exec(stmt, vals...).Error
	if err != nil {
		log.Error(err)
	}
}

func (*scannerDaoImpl) GetKnownOldest(chain string, dest interface{}) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	//stmt := fmt.Sprintf("SELECT min(`block`) FROM %s", blockTable)
	db.Table(blockTable).Select(" min(block)").Scan(dest)
}

func (*scannerDaoImpl) GetKnownOldestDone(chain string, dest interface{}) {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_block", chain)
	//stmt := fmt.Sprintf("SELECT min(`block`) FROM %s", blockTable)
	db.Table(blockTable).Select(" min(block)").Where(
		"done = true").Scan(dest)
}

func (*scannerDaoImpl) CleanOldBlock(chain string, threshold uint64) bool {
	blockTable := fmt.Sprintf("s_%s_block", chain)
	main := MainDB()

	result := main.Exec(fmt.Sprintf("DELETE FROM %s WHERE done = true AND block < %d", blockTable, threshold))
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	log.Infof("[  %v  ] clean %v old block(<%v)\n", chain, result.RowsAffected, threshold)
	return true
}

func (*scannerDaoImpl) ExistTxHash(chain, hash string) bool {
	db := MainDB()
	blockTable := fmt.Sprintf("s_%s_tx", chain)
	var count int64
	db.Table(blockTable).Where("hash = ?", hash).Count(&count)
	return count > 0
}

func (*scannerDaoImpl) NewBlockHash(tx *gorm.DB, chain string, txHash string, t time.Time) error {
	return tx.Exec(fmt.Sprintf("INSERT INTO s_%s_tx(hash, tx_time) VALUES(?,?)", chain), txHash, t.Unix()).Error
}

const (
	_30days = 30 * 24 * time.Hour
)

func (*scannerDaoImpl) Clear30DaysTx() (uint32, error) {
	main := MainDB()
	result := main.Exec("DELETE FROM s_bsc_tx WHERE tx_time <= ?", time.Now().Add(-_30days).Unix())
	if err := result.Error; err != nil {
		return 0, err
	}
	return uint32(result.RowsAffected), nil
}
