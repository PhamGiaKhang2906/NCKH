package diem

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func TongDiemRenLuyen(c *gin.Context) {
	var KiemTra struct {
		IDDiem int `json:"id_diem"`
	}

	var Mucs []models.Muc
	var DiemCongTamThoi int
	var Diem models.Diem

	if err := c.ShouldBindJSON(&KiemTra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Where("id_diem_muc = ?", KiemTra.IDDiem).Find(&Mucs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve data",
		})
		return
	}

	for _, Diem := range Mucs {
		DiemCongTamThoi += Diem.DiemMuc
	}

	if err := initializes.DB.Where("id_diem = ?", KiemTra.IDDiem).First(&Diem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve Diem data",
		})
		return
	}

	if DiemCongTamThoi > 100 {
		if err := initializes.DB.Model(&models.Diem{}).Where("id_diem = ?", KiemTra.IDDiem).Update("diem_tong", 100).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update data",
			})
			return
		}
	} else {
		if err := initializes.DB.Model(&models.Diem{}).Where("id_diem = ?", KiemTra.IDDiem).Update("diem_tong", DiemCongTamThoi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update data",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Selection and calculation updated successfully",
	})
}
