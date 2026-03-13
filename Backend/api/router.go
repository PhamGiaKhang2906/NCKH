package api

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {

	api := router.Group("/api")

	BangDiemRoutes(api)
	TieuChiRoutes(api)
	SinhVienRoutes(api)
	GiangVienRoutes(api)
	HocKyRoutes(api)
	LopSinhHoatRoutes(api)
	LoginRoutes(api)
}
