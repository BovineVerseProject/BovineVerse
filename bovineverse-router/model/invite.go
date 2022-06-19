package model

type InviteInfo struct {
	Id        uint32 `gorm:"column:id"`
	Address   string `gorm:"column:address"`
	Inviter   string `gorm:"column:inviter"`
	BlockTime int64  `gorm:"column:block_time"`
}

func (*InviteInfo) TableName() string {
	return "r_invitation"
}

type InviteRewardInfo struct {
	Id         uint32 `gorm:"column:id"`
	Address    string `gorm:"column:address"`
	Inviter    string `gorm:"column:inviter"`
	RewardType uint8  `gorm:"column:reward_type"`
	Num        uint32 `gorm:"column:num"`
	BlockTime  int64  `gorm:"column:block_time"`
}

func (*InviteRewardInfo) TableName() string {
	return "r_invite_reward"
}
