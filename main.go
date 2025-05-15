package main

import (
	"github.com/PhamGiaKhang2906/NCKH/initializes"
	"github.com/PhamGiaKhang2906/NCKH/migrate"
	"github.com/PhamGiaKhang2906/NCKH/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializes.LoadEnv()
	initializes.ConnectDataBase()
}

func main() {
	migrate.MigrateData()

	router := gin.Default()

	routers.TaiKhoanAPI(router)
	routers.DiemAPI(router)
	routers.DiemCoDinhAPI(router)
	routers.DiemHocTapAPI(router)
	routers.DiemLanAPI(router)
	routers.MucAPI(router)

	router.Run()
}
