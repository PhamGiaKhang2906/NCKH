package taikhoan

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func CapNhatTaiKhoan(c *gin.Context) {
	var taiKhoan models.TaiKhoan
	if err := c.ShouldBindJSON(&taiKhoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Model(&taiKhoan).Where("id_tai_khoan = ?", taiKhoan.IdTaiKhoan).Updates(taiKhoan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update account",
		})
		return
	}
}
