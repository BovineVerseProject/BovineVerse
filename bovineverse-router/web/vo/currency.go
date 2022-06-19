package vo

type CurrencyRate struct {
	Types uint8   `json:"types"`
	Name  string  `json:"name"`
	Rate  float64 `json:"rate"`
}
