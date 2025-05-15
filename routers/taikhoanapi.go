package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/taikhoan"
	"github.com/gin-gonic/gin"
)

func TaiKhoanAPI(router *gin.Engine) {
	router.GET("/hienthitaikhoanid", taikhoan.HienThiTaiKhoanID)
	router.GET("/hienthitoanbotaikhoan", taikhoan.HienThiToanBoTaiKhoan)
	router.POST("/taotaikhoan", taikhoan.TaoTaiKhoan)
	router.PUT("/capnhattaikhoan", taikhoan.CapNhatTaiKhoan)
	router.DELETE("/xoataikhoan", taikhoan.XoaTaiKhoan)
}
