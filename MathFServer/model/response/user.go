package response

type Login struct {
	Token string `json:"token"`
	UUID  string `json:"uuid"`
}
