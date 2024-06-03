package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已保存")
}

func UserRegister(ctx *gin.Context) {
	email := ctx.PostForm("email")
	username := ctx.PostForm("username")
	password := ctx.DefaultPostForm("password", "123456")
	passwordAgain := ctx.DefaultPostForm("password-again", "123456")
	ctx.JSON(http.StatusOK, gin.H{
		"email":         email,
		"username":      username,
		"password":      password,
		"passwordAgain": passwordAgain,
	})
}
