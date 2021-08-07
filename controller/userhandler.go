package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context)  {
	//name := ctx.PostForm("name")
	pass := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg":"Phone num must be 11 numbers", "code":422})
		return
	}

	if len(pass) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg":"password's length must be greater than 6", "code":422})
		return

	}

	//utils.
	//ctx.JSON(utils.NewSucc("Register success", gin.H{
	//	"msg":"Register success!",
	//}))
}
