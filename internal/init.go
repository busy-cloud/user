package internal

import (
	"encoding/gob"
	"github.com/busy-cloud/boat/db"
)

func init() {
	gob.Register(User{})

	db.Register(new(User), new(Password))
}
