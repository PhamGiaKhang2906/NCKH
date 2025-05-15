package taikhoan

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func XoaTaiKhoan(c *gin.Context) {
	var taiKhoan models.TaiKhoan
	if err := c.ShouldBindJSON(&taiKhoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Delete(&models.TaiKhoan{}, taiKhoan.IdTaiKhoan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete account",
		})
		return
	}
}
