package model

type CurrencyInfo struct {
	Id          uint32  `gorm:"column:id"`
	Types       uint8   `gorm:"column:types"`
	Name        string  `gorm:"column:name"`
	Address     string  `gorm:"column:address"`
	Pair        string  `gorm:"column:pair"`
	Decimal     uint32  `gorm:"column:decimal"`
	Default     float64 `gorm:"column:default"`
	DynamicRate float64 `gorm:"column:dynamic_rate"`
}

func (*CurrencyInfo) TableName() string {
	return "s_currency"
}
