package taikhoan

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func HienThiTaiKhoanID(c *gin.Context) {
	var taiKhoan models.TaiKhoan
	if err := c.ShouldBindJSON(&taiKhoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Where("id_tai_khoan = ?", taiKhoan.IdTaiKhoan).First(&taiKhoan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Account not found",
		})
		return
	}

	c.JSON(http.StatusOK, taiKhoan)
}

func HienThiToanBoTaiKhoan(c *gin.Context) {
	var taiKhoans []models.TaiKhoan

	if err := initializes.DB.Find(&taiKhoans).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Failed to show accounts",
		})
		return
	}

	c.JSON(http.StatusOK, taiKhoans)
}
