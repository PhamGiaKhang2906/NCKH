package models

type Muc struct {
	IdMuc          int    `gorm:"primaryKey" json:"id_muc"`
	DiemMucToiDa   int    `json:"diem_muc_toi_da"`
	DiemMuc        int    `json:"diem_muc"`
	TongDiemCoDinh int    `json:"tong_diem_co_dinh"`
	TongDiemSoLan  int    `json:"tong_diem_so_lan"`
	TongDiemHocTap int    `json:"tong_diem_hoc_tap"`
	TenMuc         string `json:"ten_muc"`
	IdDiemMuc      int    `json:"id_diem"`

	MucDiem Diem `gorm:"foreignKey:IdDiemMuc;references:IdDiem" json:"muc_diem"`

	DiemCoDinh []DiemCoDinh `gorm:"foreignKey:IdMucCD;references:IdMuc" json:"diem_co_dinh"`
	DiemHocTap []DiemHocTap `gorm:"foreignKey:IdMucHT;references:IdMuc" json:"diem_hoc_tap"`
	DiemLan    []DiemLan    `gorm:"foreignKey:IdMucL;references:IdMuc" json:"diem_lan"`
}
