package muc

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func TongDiemMuc(c *gin.Context) {
	var Muc models.Muc

	if err := c.ShouldBindJSON(&Muc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Where("id_muc = ?", Muc.IdMuc).First(&Muc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve Muc data",
		})
		return
	}

	Muc.DiemMuc = Muc.TongDiemCoDinh + Muc.TongDiemSoLan + Muc.TongDiemHocTap

	if Muc.DiemMuc < Muc.DiemMucToiDa {
		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", Muc.IdMuc).Updates(Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	} else {
		Muc.DiemMuc = Muc.DiemMucToiDa
		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", Muc.IdMuc).Updates(Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	}
}
