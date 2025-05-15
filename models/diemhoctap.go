package models

type DiemHocTap struct {
	IdDiemHocTap int     `gorm:"primaryKey" json:"id_diem_hoc_tap"`
	TrungBinhKy  float64 `json:"trung_binh_ky"`
	DiemRenLuyen int     `json:"diem_ren_luyen"`
	IdMucHT      int     `json:"id_muc_hoc_tap"`

	DiemMuc Muc `gorm:"foreignKey:IdMucHT;references:IdMuc;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"diem_muc"`
}
