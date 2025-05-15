package diemcodinh

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func Chon(c *gin.Context) {
	var Chon struct {
		IDDiemCD   int  `json:"id_diem_cd"`
		GiaTriChon bool `json:"chon"`
	}

	if err := c.ShouldBindJSON(&Chon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if Chon.GiaTriChon {
		var DiemCoDinh models.DiemCoDinh
		var Muc models.Muc

		if err := initializes.DB.Model(&models.DiemCoDinh{}).Where("id_diem_co_dinh = ?", Chon.IDDiemCD).Update("chon", 1).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update selection",
			})
			return
		}

		if err := initializes.DB.Where("id_diem_co_dinh = ?", Chon.IDDiemCD).First(&DiemCoDinh).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve DiemCoDinh data",
			})
			return
		}

		if err := initializes.DB.Where("id_muc = ?", DiemCoDinh.IdMucCD).First(&Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve Muc data",
			})
			return
		}

		Muc.TongDiemCoDinh += DiemCoDinh.DiemCong

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", Muc.IdMuc).Updates(Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}

	} else {
		var DiemCoDinh models.DiemCoDinh
		var Muc models.Muc

		if err := initializes.DB.Model(&models.DiemCoDinh{}).Where("id_diem_co_dinh = ?", Chon.IDDiemCD).Update("chon", 0).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update selection",
			})
			return
		}

		if err := initializes.DB.Where("id_diem_co_dinh = ?", Chon.IDDiemCD).First(&DiemCoDinh).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve DiemCoDinh data",
			})
			return
		}

		if err := initializes.DB.Where("id_muc = ?", DiemCoDinh.IdMucCD).First(&Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve Muc data",
			})
			return
		}

		Muc.TongDiemCoDinh -= DiemCoDinh.DiemCong

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", Muc.IdMuc).Updates(Muc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	}
}
