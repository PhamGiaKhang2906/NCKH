package routers

import (
	"github.com/PhamGiaKhang2906/NCKH/controllers/diem/muc/diemsolan"
	"github.com/gin-gonic/gin"
)

func DiemLanAPI(router *gin.Engine) {
	router.PUT("/chonsolan", diemsolan.ChonSoLan)
}
