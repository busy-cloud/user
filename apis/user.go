package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/user"
)

const BasePath = "app/user/api/"

func init() {
	api.Register("GET", BasePath+"user/list", curd.ApiList[user.User]())
}
