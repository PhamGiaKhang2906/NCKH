package api

import (
	"Backend/controller/giangvien"

	"github.com/gin-gonic/gin"
)

func GiangVienRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xemgiangvien/:magiangvien", giangvien.XemGiangVien)
	router.GET("/api/xemtatcagiangvien", giangvien.XemTatCaGiangVien)

	// POST*
	router.POST("/api/taogiangvien", giangvien.TaoGiangVien)

	// PUT*
	router.PUT("/api/suagiangvien", giangvien.SuaGiangVien)

	// DELETE*
	router.DELETE("/api/xoagiangvien/:magiangvien", giangvien.XoaGiangVien)
}
