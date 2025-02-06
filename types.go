package user

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(&User{}, &Password{})
}

type User struct {
	Id        string `json:"id" xorm:"pk"`
	Name      string
	Email     string
	CreatedAt time.Time
}

type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
