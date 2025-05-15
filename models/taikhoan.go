package models

type TaiKhoan struct {
	IdTaiKhoan  int    `gorm:"primaryKey" json:"id_tai_khoan"`
	TenTaiKhoan string `json:"ten_tai_khoan"`
	MatKhau     string `json:"mat_khau"`
	Email       string `json:"email"`

	TaiKhoanSinhVien *SinhVien `gorm:"foreignKey:IdTaiKhoanSinhVien;references:IdTaiKhoan" json:"tai_khoan_sinh_vien"`
}
