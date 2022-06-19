package model

type BreedInfo struct {
	Id       uint32  `gorm:"column:id"`
	TokenId  uint32  `gorm:"column:token_id"`
	Price    float64 `gorm:"column:price"`
	Seller   string  `gorm:"column:seller"`
	CreateAt int64   `gorm:"column:create_at"`
}

func (*BreedInfo) TableName() string {
	return "u_breed_center"
}

type BreedCattleInfo struct {
	Id         uint32  `gorm:"column:id"`
	TokenId    uint32  `gorm:"column:token_id"`
	Price      float64 `gorm:"column:price"`
	Seller     string  `gorm:"column:seller"`
	CreateAt   int64   `gorm:"column:create_at"`
	BreedCount uint32  `gorm:"-"`

	CattleInfo
}
