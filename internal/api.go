package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
)

func init() {

	api.RegisterUnAuthorized("POST", "user/login", login)
	//api.RegisterUnAuthorized("POST", "login", login)

	api.RegisterUnAuthorized("POST", "user/auth", auth)
	//api.RegisterUnAuthorized("GET", "auth", auth)

	api.Register("GET", "user/logout", logout)
	//api.Register("GET", "logout", logout)

	api.Register("POST", "user/password", password)
	//api.Register("POST", "password", password)

	api.Register("GET", "user/me", userMe)
	//api.Register("GET", "me", userMe)

	api.Register("POST", "user/count", curd.ApiCount[User]())

	api.Register("POST", "user/search", curd.ApiSearch[User]())

	api.Register("GET", "user/list", curd.ApiList[User]())

	api.Register("POST", "user/create", curd.ApiCreateHook[User](curd.GenerateID[User](), nil))

	api.Register("GET", "user/:id", curd.ApiGet[User]())

	api.Register("POST", "user/:id", curd.ApiUpdate[User]("id", "name", "email", "cellphone", "admin", "disabled"))

	api.Register("GET", "user/:id/delete", curd.ApiDeleteHook[User](nil, nil))

	api.Register("POST", "user/:id/password", userPassword)

	api.Register("GET", "user/:id/enable", curd.ApiDisableHook[User](false, nil, nil))

	api.Register("GET", "user/:id/disable", curd.ApiDisableHook[User](true, nil, nil))
}

func userMe(ctx *gin.Context) {
	id := ctx.GetString("user")
	var user User
	has, err := db.Engine().ID(id).Get(&user)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if !has {
		api.Fail(ctx, "用户不存在")
		return
	}
	api.OK(ctx, &user)
}

type userPasswordObj struct {
	Password string `json:"password"`
}

func userPassword(ctx *gin.Context) {

	var obj userPasswordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	var p Password
	p.Id = ctx.Param("id")
	p.Password = md5hash(obj.Password)

	_, _ = db.Engine().ID(p.Id).Delete(new(Password)) //不管有没有都删掉
	_, err := db.Engine().InsertOne(&p)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
