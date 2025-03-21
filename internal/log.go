package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {

	api.Register("POST", "user/log/count", curd.ApiCount[UserLog]())

	api.Register("POST", "user/log/search", curd.ApiSearch[UserLog]())

	api.Register("GET", "user/log/list", curd.ApiList[UserLog]())

	api.Register("POST", "user/log/create", curd.ApiCreate[UserLog]())

	api.Register("GET", "user/log/:id", curd.ApiGet[UserLog]())

	api.Register("GET", "user/log/:id/delete", curd.ApiDelete[UserLog]())

}
