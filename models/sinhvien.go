package models

type SinhVien struct {
	IdSinhVien         int    `gorm:"primaryKey" json:"id_sinh_vien"`
	TenSinhVien        string `json:"ten_sinh_vien"`
	Lop                string `json:"lop"`
	Khoa               string `json:"khoa"`
	IdTaiKhoanSinhVien int    `json:"id_tai_khoan_sinh_vien"`

	DiemSinhVien     *Diem     `gorm:"foreignKey:IdDiemSinhVien;references:IdSinhVien" json:"diem_sinh_vien"`
	TaiKhoanSinhVien *TaiKhoan `gorm:"foreignKey:IdTaiKhoanSinhVien;references:IdTaiKhoan" json:"tai_khoan_sinh_vien"`
}
