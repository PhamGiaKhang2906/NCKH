package api

import (
	"Backend/controller"
	"Backend/controller/login"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.RouterGroup) {
	//GET*

	//POST*
	router.POST("/api/login", controller.Login)
	router.POST("/api/sauloginsinhvien", login.SauLoginSinhVien)
	router.POST("/api/saulogingiangvien", login.SauLoginGiangVien)

	//PUT*

	//DELETE*
}
