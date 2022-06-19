package model

type Block struct {
	Block      uint64 `gorm:"column:block; primaryKey;"`
	FailTimes  uint8  `gorm:"column:fail_times; default:0;" json:"fail_times"`
	Done       bool   `gorm:"column:done; NOT NULL; default:0;" json:"done"`
	Time       int64  `gorm:"column:time; type:int8; default:0;" json:"time"`
	DispatchTs int64  `gorm:"column:dispatch_ts; type:int8; NOT NULL; default:0;" json:"dispatch_ts"`
}

func (*Block) TableName() string {
	return "s_bsc_block"
}

type BlockTx struct {
	Id     int64  `gorm:"column:id; primaryKey;"`
	Hash   string `gorm:"column:hash; type:varchar(255); not null; uniqueIndex;"`
	TxTime int64  `gorm:"column:tx_time; type:int8; NOT NULL; default:0;" json:"tx_time"`
}

func (*BlockTx) TableName() string {
	return "s_bsc_tx"
}
