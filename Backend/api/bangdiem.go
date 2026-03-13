package api

import (
	"Backend/controller/bangdiem"

	"github.com/gin-gonic/gin"
)

func BangDiemRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xembangdiem", bangdiem.XemBangDiem)
	router.GET("/api/xemdanhsachbangdiemsinhvien/:malopsinhhoat/:mahocky", bangdiem.XemDanhSachBangDiemSinhVien)
	router.GET("/api/xemdanhsachbangdiemsinhvientheolop/:magiangvien/:mahocky", bangdiem.XemDanhSachBangDiemSinhVienTheoLop)
	router.GET("/api/xemdanhsachbangdiemsinhvientheokhoa/:machuyenvien/:mahocky", bangdiem.XemDanhSachBangDiemSinhVienTheoKhoa)
	router.GET("/api/xemdiemdachamquacacnam/:masinhvien", bangdiem.XemDiemDaChamQuaCacNam)
	router.GET("/api/xemtrangthaibangdiem/:mabangdiem", bangdiem.XemTrangThaiBangDiem)

	// POST*
	router.POST("/api/taobangdiem", bangdiem.TaoBangDiem)
	router.POST("/api/saochepbangdiem", bangdiem.SaoChepBangDiem)
	router.POST("/api/phatbangdiem", bangdiem.PhatBangDiem)

	// PUT*
	router.PUT("/api/thaydoitrangthai", bangdiem.ThayDoiTrangThaiSinhVien)
	router.PUT("/api/thaydoitrangthaigiangvien", bangdiem.ThayDoiTrangThaiGiangVien)
	router.PUT("/api/thaydoitrangthaichuyenvien", bangdiem.ThayDoiTrangThaiChuyenVien)

	// DELETE*
	router.DELETE("/api/xoabangdiem/:mabangdiem", bangdiem.XoaBangDiem)
}
