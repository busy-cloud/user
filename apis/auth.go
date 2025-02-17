package user

import (
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
)

func auth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	var user User
	has, err := db.Engine.Where("id=?", username).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	if !has {
		curd.Fail(ctx, "找不到用户")
		return
	}

	if user.Disabled {
		curd.Fail(ctx, "用户已禁用")
		return
	}

	var obj Password
	has, err = db.Engine.ID(user.Id).Get(&obj)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := "123456"
		obj.Password = md5hash(dp)
	}

	if obj.Password != password {
		curd.Fail(ctx, "密码错误")
		return
	}

	//生成Token
	token, err := web.JwtGenerate(user.Id, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, gin.H{
		token: token,
	})
}
