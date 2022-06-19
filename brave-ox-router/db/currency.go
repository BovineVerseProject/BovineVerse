package db

var (
	CurrencyDao currencyDao = &currencyDaoImpl{}
)

type currencyDao interface {
	UpdateRate(types uint8, price float64) error
}

type currencyDaoImpl struct{}

func (*currencyDaoImpl) UpdateRate(types uint8, price float64) error {
	return MainDB().Table("s_currency").
		Where("types = ? AND dynamic_rate != ?", types, price).UpdateColumn("dynamic_rate", price).Error
}
