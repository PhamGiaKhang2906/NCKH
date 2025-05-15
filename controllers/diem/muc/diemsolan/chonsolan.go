package diemsolan

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func ChonSoLan(c *gin.Context) {
	var CheckSoLan struct {
		IDDiemSL int `json:"id_diem_so_lan"`
		SoLan    int `json:"so_lan"`
	}

	if err := c.ShouldBindJSON(&CheckSoLan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if CheckSoLan.SoLan != 0 {
		var DiemLan models.DiemLan
		var Muc models.Muc

		if err := initializes.DB.Where("id_diem_lan = ?", CheckSoLan.IDDiemSL).First(&DiemLan).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve DiemLan data",
			})
			return
		}

		DiemLan.SoLan = CheckSoLan.SoLan

		DiemLan.DiemCuoiCung = CheckSoLan.SoLan * DiemLan.DiemMotLan

		if DiemLan.DiemGhiNho == 0 {
			if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemLan.IdMucL).First(&Muc).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to retrieve Muc data",
				})
				return
			}

			Muc.TongDiemSoLan += DiemLan.DiemCuoiCung

			if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemLan.IdMucL).Updates(Muc).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to update Muc data",
				})
				return
			}

			DiemLan.DiemGhiNho = DiemLan.DiemCuoiCung

			if err := initializes.DB.Model(&models.DiemLan{}).Where("id_diem_lan = ?", CheckSoLan.IDDiemSL).Updates(DiemLan).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to update DiemLan data",
				})
				return
			}
		} else {
			if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemLan.IdMucL).First(&Muc).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to retrieve Muc data",
				})
				return
			}

			Muc.TongDiemSoLan -= DiemLan.DiemGhiNho
			Muc.TongDiemSoLan += DiemLan.DiemCuoiCung

			if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemLan.IdMucL).Updates(Muc).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to update Muc data",
				})
				return
			}

			DiemLan.DiemGhiNho = DiemLan.DiemCuoiCung

			if err := initializes.DB.Model(&models.DiemLan{}).Where("id_diem_lan = ?", CheckSoLan.IDDiemSL).Updates(DiemLan).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to update DiemLan data",
				})
				return
			}
		}
	}
}
