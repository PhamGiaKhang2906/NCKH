package migrate

import (
	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/models"
)

func MigrateData() {
	initializes.DB.AutoMigrate(&models.SinhVien{})
	initializes.DB.AutoMigrate(&models.TaiKhoan{})
	initializes.DB.AutoMigrate(&models.Diem{})
	initializes.DB.AutoMigrate(&models.Muc{})
	initializes.DB.AutoMigrate(&models.DiemHocTap{})
	initializes.DB.AutoMigrate(&models.DiemCoDinh{})
	initializes.DB.AutoMigrate(&models.DiemLan{})
}
