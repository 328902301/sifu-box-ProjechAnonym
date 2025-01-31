package route

import (
	"net/http"
	"sifu-box/control"
	"sifu-box/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SettingLogin(api *gin.RouterGroup, user *models.User, logger *zap.Logger){
	api.POST("/login/:user", func(ctx *gin.Context){
		switch ctx.Param("user") {
		case "visitor":
			code := ctx.PostForm("code")
			if code == user.Code {
				token, err := control.Login(false, user.PrivateKey, logger)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(),})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"message": struct{
					JWT string `json:"jwt"`
					Admin bool `json:"admin"`
				}{JWT: token, Admin: false}})
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "密钥错误",})
			
		case "admin":
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")
			if username == user.Username && password == user.Password {
				token, err := control.Login(true, user.PrivateKey, logger)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(),})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"message": 
					struct{
						JWT string `json:"jwt"`
						Admin bool `json:"admin"`
					}{JWT: token, Admin: true}})
				return
			} 
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
			
		default:
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "用户参数不正确",})	
		}
	})
	api.GET("/verify", func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		token, admin, err := control.Verify(authorization, user.PrivateKey, logger)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": struct{
			JWT string `json:"jwt"`
			Admin bool `json:"admin"`
		}{JWT: token, Admin: admin}})
	})
}