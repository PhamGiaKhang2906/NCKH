package api

import (
	"Backend/controller/tieuchi"

	"github.com/gin-gonic/gin"
)

func TieuChiRoutes(router *gin.RouterGroup) {
	// GET*
	router.GET("/api/xemtieuchi/:mabangdiem", tieuchi.XemTieuChi)
	router.GET("/api/xemtieuchicham/:masinhvien/:mahocky", tieuchi.XemTieuChiCham)

	// POST*
	router.POST("/api/taotieuchi", tieuchi.TaoTieuChi)

	// PUT*
	router.PUT("/api/suatieuchi", tieuchi.SuaTieuChi)
	router.PUT("/api/chamdiem", tieuchi.ChamDiem)
	router.PUT("/api/saochepdiem", tieuchi.SaoChepDiem)
	router.PUT("/api/saocheptoanbodiem", tieuchi.SaoChepToanBoDiem)

	// DELETE*
}
