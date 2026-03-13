package api

import (
	"Backend/controller/lopsinhhoat"

	"github.com/gin-gonic/gin"
)

func LopSinhHoatRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xemlopsinhhoat/:malopsinhhoat", lopsinhhoat.XemLopSinhHoat)
	router.GET("/api/xemtatcalopsinhhoat", lopsinhhoat.XemTatCaLopSinhHoat)

	// POST*
	router.POST("/api/taolopsinhhoat", lopsinhhoat.TaoLopSinhHoat)
	router.POST("/api/taolopsinhhoathocky", lopsinhhoat.TaoLopSinhHoatHocKy)
	router.POST("/api/taolopsinhhoatsinhvien", lopsinhhoat.TaoLopSinhHoatSinhVien)

	// PUT*
	router.PUT("/api/sualopsinhhoat", lopsinhhoat.SuaLopSinhHoat)

	// DELETE*
	router.DELETE("/api/xoalopsinhhoat/:malopsinhhoat", lopsinhhoat.XoaLopSinhHoat)
}
