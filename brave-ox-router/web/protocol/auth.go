package protocol

type AuthRequest struct {
	Address string `json:"address" binding:"required,len=42"`
	Sig     string `json:"sig" binding:"required"`
}

type AuthResponse struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}
