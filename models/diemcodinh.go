package models

type DiemCoDinh struct {
	IdDiemCoDinh  int    `gorm:"primaryKey" json:"id_diem_co_dinh"`
	TenDiemCoDinh string `json:"ten_diem_co_dinh"`
	DiemCong      int    `json:"diem_cong"`
	Chon          bool   `json:"chon" gorm:"default:0"`
	IdMucCD       int    `json:"id_muc_co_dinh"`

	DiemMuc Muc `gorm:"foreignKey:IdMucCD;references:IdMuc;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"diem_muc"`
}
