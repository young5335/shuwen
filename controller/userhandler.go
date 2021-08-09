package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(ctx *gin.Context)  {
	name := ctx.PostForm("name")
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

	if len(name) == 0 {
		name = randStr(10)
	}
	log.Println(name)

	//utils.
	ctx.JSON(utils.NewSucc("Register success", gin.H{
		"msg":"Register success!",
	}))
}

func randStr(n int) string {
	a := byte('a')
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	chars := genChars(a)
	for i := 0; i < n; i++ {
		result = append(result, chars[rand.Intn(26)])
	}
	return string(result)
}
func genChars(startChar byte) string {
	result :=[]byte{}
	for i := 0; i < 26; i++ {
		result = append(result, startChar + byte(i))
	}
	return string(result)
}