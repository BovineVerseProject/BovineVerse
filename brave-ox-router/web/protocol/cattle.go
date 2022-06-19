package protocol

type CattleVO struct {
	Id        uint32 `gorm:"column:id" json:"id" `
	Class     uint8  `gorm:"column:class" json:"class"`
	Gender    uint8  `gorm:"column:gender" json:"gender"`
	GenderSeq uint32 `gorm:"column:gender_seq" json:"gsq"`
	Life      uint32 `gorm:"column:life" json:"life"`
	Growth    uint32 `gorm:"column:growth" json:"growth"`
	Energy    uint32 `gorm:"column:energy" json:"energy"`
	IsAdult   bool   `gorm:"column:is_adult" json:"isAdult"`
	Star      uint8  `gorm:"column:star" json:"star"`
	Attack    uint32 `gorm:"column:attack" json:"attack"`
	Stamina   uint32 `gorm:"column:stamina" json:"stamina"`
	Defense   uint32 `gorm:"column:defense" json:"defense"`
	Milk      uint32 `gorm:"column:milk" json:"milk"`
	MilkRate  uint32 `gorm:"column:milk_rate" json:"milkRate"`
	Image     string `gorm:"column:image" json:"image" json:"image"`
	DeadAt    int64  `gorm:"column:dead_at" json:"-"`
}

type CattleListRequest struct {
	Ids []uint32 `json:"ids" binding:"required,lte=60"`
}
