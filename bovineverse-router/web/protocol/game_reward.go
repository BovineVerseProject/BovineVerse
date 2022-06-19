package protocol

type ClaimGameRewardRequest struct {
	Address     string   `json:"address" binding:"required,len=42"`
	Types       uint8    `json:"types" binding:"required,oneof=1 2 3 4 5"`
	TypeSeq     uint32   `json:"typeSeq" binding:"required,gte=1"`
	RewardTypes []int32  `json:"rewardTypes" binding:"required,lte=50"`
	ItemId      []int32  `json:"itemId" binding:"required,lte=50"`
	Count       []string `json:"count" binding:"required,lte=50"`
}

type ClaimGameRewardResponse struct {
	Id       uint64 `json:"id"`
	Finished bool   `json:"finished"`
	Types    uint8  `json:"types"`
	TypeSeq  uint32 `json:"typeSeq"`
	R        string `json:"r"`
	S        string `json:"s"`
	V        uint8  `json:"v"`
}

type ClaimRewardStatusRequest struct {
	Ids []uint64 `json:"ids" binding:"required,lte=50"`
}

type ClaimRewardStatus struct {
	Id       uint64 `json:"id"`
	Finished bool   `json:"finished"`
}
