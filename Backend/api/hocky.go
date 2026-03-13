package api

import (
	"Backend/controller/hocky"

	"github.com/gin-gonic/gin"
)

func HocKyRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xemhocky", hocky.XemHocKy)
	router.GET("/api/xemdanhsachhocky/:manguoidung/:type", hocky.XemDanhSachHocKy)

	// POST*

	// PUT*
	router.PUT("/api/doihocky", hocky.DoiHocKy)

	// DELETE*
	router.DELETE("/api/xoahocky/:mahocky", hocky.XoaHocKy)

}
