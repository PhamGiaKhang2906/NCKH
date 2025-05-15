package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/diem/muc"
	"github.com/gin-gonic/gin"
)

func MucAPI(router *gin.Engine) {
	router.PUT("/tongdiemmuc", muc.TongDiemMuc)
}
