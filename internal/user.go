package internal

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(new(User), new(Password), new(UserLog), new(Role))
}

// User 用户
type User struct {
	Id        string    `json:"id" xorm:"pk"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}

// Role 角色
type Role struct {
	Id         string    `json:"id" xorm:"pk"`
	Name       string    `json:"name,omitempty"`
	Privileges []string  `json:"privileges,omitempty" xorm:"LONGTEXT"`
	Disabled   bool      `json:"disabled,omitempty"`
	Created    time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}

type UserLog struct {
	Id      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	Action  string    `json:"action,omitempty"`
	Client  string    `json:"client,omitempty"`
	Ip      string    `json:"ip,omitempty"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}
