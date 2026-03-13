package api

import (
	"Backend/controller/sinhvien"

	"github.com/gin-gonic/gin"
)

func SinhVienRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xemsinhvien/:masinhvien", sinhvien.XemSinhVien)
	router.GET("/api/xemtatcasinhvien", sinhvien.XemTatCaSinhVien)

	// POST*
	router.POST("/api/taosinhvien", sinhvien.TaoSinhVien)

	// PUT*
	router.PUT("/api/suasinhvien", sinhvien.SuaSinhVien)

	// DELETE*
	router.DELETE("/api/xoasinhvien/:masinhvien", sinhvien.XoaSinhVien)
}
