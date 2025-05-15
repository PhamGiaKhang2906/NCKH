package models

type Diem struct {
	IdDiem         int `gorm:"primaryKey" json:"id_diem"`
	DiemTong       int `json:"diem_tong"`
	IdDiemSinhVien int `json:"id_diem_sinh_vien"`

	DiemSinhVien *SinhVien `gorm:"foreignKey:IdDiemSinhVien;references:IdSinhVien" json:"diem_sinh_vien"`

	MucDiem []Muc `gorm:"foreignKey:IdDiemMuc;references:IdDiem" json:"muc_diem"`
}
