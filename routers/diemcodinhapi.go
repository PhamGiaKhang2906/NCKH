package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/diem/muc/diemcodinh"
	"github.com/gin-gonic/gin"
)

func DiemCoDinhAPI(router *gin.Engine) {
	router.PUT("/chon", diemcodinh.Chon)
}
