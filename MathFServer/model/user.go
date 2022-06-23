package model

type User struct {
	UUID      string   `json:"uuid" gorm:"primary_key"`
	UserName  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Remaining int      `json:"remaining"`
	VipLevel  VipLevel `json:"vip_level"`
}

type VipLevel int

const (
	Lv0 VipLevel = VipLevel(iota)
	Lv1
	Lv2
)
