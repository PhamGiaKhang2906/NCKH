package models

type DiemLan struct {
	IdDiemLan    int    `gorm:"primaryKey" json:"id_diem_lan"`
	TenDiemLan   string `json:"ten_diem_lan"`
	DiemMotLan   int    `json:"diem_mot_lan"`
	SoLan        int    `json:"so_lan"`
	DiemCuoiCung int    `json:"diem_cuoi_cung"`
	DiemGhiNho   int    `json:"diem_ghi_nho"`
	IdMucL       int    `json:"id_muc_lan"`

	DiemMuc Muc `gorm:"foreignKey:IdMucL;references:IdMuc;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"diem_muc"`
}
