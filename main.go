package main

import (
	"home-loans/auth"
	"home-loans/config"
	"home-loans/handler"
	"home-loans/kelengkapan"
	"home-loans/middleware"
	"home-loans/pengajuan"
	"home-loans/user"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	config := config.Init()
	dsn := config.DBDSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//migration
	db.Migrator().AutoMigrate(&user.User{}, &pengajuan.Pengajuan{}, &kelengkapan.Kelengkapan{})

	//MIDDLEWARE
	userMiddleware := middleware.UserMiddleware
	staffMiddleware := middleware.StaffMiddleware

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//PENGAJUAN
	pengajuanRepository := pengajuan.NewRepository(db)
	pengajuanService := pengajuan.NewService(pengajuanRepository)
	pengajuanHandler := handler.NewPengajuanHandler(pengajuanService, authService)

	//KELENGKAPAN
	kelengkapanRepository := kelengkapan.NewRepository(db)
	kelengkapanService := kelengkapan.NewService(kelengkapanRepository)
	kelengkapanHandler := handler.NewKelengkapanHandler(kelengkapanService, authService)

	router := gin.Default()
	router.Use(cors.Default())

	//grouping
	api := router.Group("/api/v1")

	//publish file
	api.Static("bukti_ktp_files/", "./bukti_ktp_files")
	api.Static("bukti_slip_gaji_files/", "./bukti_slip_gaji_files")
	api.Static("dokumen_pendukung_files/", "./dokumen_pendukung_files")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	//pengajuan endpoint user
	api.GET("/pengajuan-user", userMiddleware(authService, userService), pengajuanHandler.GetPengajuanUser)
	api.POST("/pengajuan", userMiddleware(authService, userService), pengajuanHandler.CreatePengajuan)
	api.PUT("/pengajuan/bukti-ktp", userMiddleware(authService, userService), pengajuanHandler.UploadBuktiKtp)
	api.PUT("/pengajuan/bukti-slip-gaji", userMiddleware(authService, userService), pengajuanHandler.UploadBuktiSlipGaji)
	api.DELETE("/pengajuan", userMiddleware(authService, userService), pengajuanHandler.DeletePengajuan)

	//pengajuan endpoint staff
	api.GET("/pengajuan", staffMiddleware(authService, userService), pengajuanHandler.GetPengajuans)
	api.GET("/pengajuan/check-recommendation", staffMiddleware(authService, userService), pengajuanHandler.CheckRecommendation)
	api.PUT("/pengajuan/status/:user_id", staffMiddleware(authService, userService), pengajuanHandler.UpdatePengajuanStatus)

	//kelengkapan endpoint user
	api.GET("/kelengkapan-user", userMiddleware(authService, userService), kelengkapanHandler.GetKelengkapanUser)
	api.POST("/kelengkapan", userMiddleware(authService, userService), kelengkapanHandler.CreateKelengkapan)
	api.PUT("/kelengkapan/dokumen-pendukung", userMiddleware(authService, userService), kelengkapanHandler.UploadDokumenPendukung)
	api.DELETE("/kelengkapan", userMiddleware(authService, userService), kelengkapanHandler.DeleteKelengkapan)

	//kelengkapan endpoint staff
	api.GET("/kelengkapan", staffMiddleware(authService, userService), kelengkapanHandler.GetKelengkapans)
	api.PUT("/kelengkapan/status/:user_id", staffMiddleware(authService, userService), kelengkapanHandler.UpdateKelengkapanStatus)

	router.Run()

}
