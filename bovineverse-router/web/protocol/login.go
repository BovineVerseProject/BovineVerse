package protocol

type LoginResult struct {
	User  string `json:"user"`
	Token string `json:"token"`
}
