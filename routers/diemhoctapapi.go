package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/diem/muc/diemhoctap"
	"github.com/gin-gonic/gin"
)

func DiemHocTapAPI(router *gin.Engine) {
	router.PUT("/ketquahoctap", diemhoctap.KetQuaHocTap)
}
