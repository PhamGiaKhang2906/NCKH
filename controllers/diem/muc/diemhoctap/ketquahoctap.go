package diemhoctap

import (
	"net/http"

	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
	"github.com/gin-gonic/gin"
)

func KetQuaHocTap(c *gin.Context) {
	var DiemHT struct {
		IdDiemMucHT int     `json:"id_diem_muc_ht"`
		Diem        float64 `json:"diem"`
	}
	var DiemHocTap models.DiemHocTap

	if err := c.ShouldBindJSON(&DiemHT); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := initializes.DB.Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).First(&DiemHocTap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve DiemHocTap data",
		})
		return
	}

	switch {
	case DiemHT.Diem <= 4.0 && DiemHT.Diem >= 3.6:
		DiemHocTap.DiemRenLuyen = 11
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 11).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}

	case DiemHT.Diem >= 3.2 && DiemHT.Diem < 3.6:
		DiemHocTap.DiemRenLuyen = 9
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 9).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	case DiemHT.Diem >= 2.5 && DiemHT.Diem < 3.2:
		DiemHocTap.DiemRenLuyen = 7
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 7).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	case DiemHT.Diem >= 2.0 && DiemHT.Diem < 2.5:
		DiemHocTap.DiemRenLuyen = 5
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 5).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	case DiemHT.Diem >= 1.2 && DiemHT.Diem < 2.0:
		DiemHocTap.DiemRenLuyen = 3
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 3).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	case DiemHT.Diem >= 0.0 && DiemHT.Diem < 1.2:
		DiemHocTap.DiemRenLuyen = 0
		DiemHocTap.TrungBinhKy = DiemHT.Diem

		if err := initializes.DB.Model(&models.DiemHocTap{}).Where("id_diem_hoc_tap = ?", DiemHT.IdDiemMucHT).Updates(DiemHocTap).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update DiemHocTap data",
			})
			return
		}

		if err := initializes.DB.Model(&models.Muc{}).Where("id_muc = ?", DiemHocTap.IdMucHT).Update("tong_diem_hoc_tap", 0).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update Muc data",
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid score",
		})
		return
	}
}
