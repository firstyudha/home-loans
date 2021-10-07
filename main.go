package main

import (
	"home-loans/auth"
	"home-loans/config"
	"home-loans/handler"
	"home-loans/helper"
	"home-loans/kelengkapan"
	"home-loans/pengajuan"
	"home-loans/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

func userMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Split by space " "
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate token
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//context currentUser
		c.Set("currentUser", user)

	}

}

func staffMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Split by space " "
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate token
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if user.LoginAs != 2 { //must be officer/staff
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//context currentUser
		c.Set("currentUser", user)

	}

}
