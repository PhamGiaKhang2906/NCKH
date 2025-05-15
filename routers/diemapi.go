package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/diem"
	"github.com/gin-gonic/gin"
)

func DiemAPI(router *gin.Engine) {
	router.PUT("/tongdiemrenluyen", diem.TongDiemRenLuyen)
}
