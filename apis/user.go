package user

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(new(User), new(Password))
}

// User 用户
type User struct {
	Id        string    `json:"id" xorm:"pk"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
