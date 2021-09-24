package main

import (
	"home-loans/auth"
	"home-loans/handler"
	"home-loans/helper"
	"home-loans/kelengkapan"
	"home-loans/pengajuan"
	"home-loans/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-homeloans?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

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

	//publish file
	router.Static("bukti_ktp_files/", "./bukti_ktp_files")
	router.Static("bukti_slip_gaji_files/", "./bukti_slip_gaji_files")
	router.Static("dokumen_pendukung_files/", "./dokumen_pendukung_files")

	//grouping
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	//pengajuan endpoint
	api.GET("/pengajuan", authMiddleware(authService, userService), pengajuanHandler.GetPengajuans)
	api.POST("/pengajuan", authMiddleware(authService, userService), pengajuanHandler.CreatePengajuan)
	api.PUT("/pengajuan/bukti-ktp", authMiddleware(authService, userService), pengajuanHandler.UploadBuktiKtp)
	api.PUT("/pengajuan/bukti-slip-gaji", authMiddleware(authService, userService), pengajuanHandler.UploadBuktiSlipGaji)

	//kelengkapan endpoint
	api.GET("/kelengkapan", authMiddleware(authService, userService), kelengkapanHandler.GetKelengkapans)
	api.POST("/kelengkapan", authMiddleware(authService, userService), kelengkapanHandler.CreateKelengkapan)
	api.PUT("/kelengkapan/dokumen-pendukung/:pengajuan_id", authMiddleware(authService, userService), kelengkapanHandler.UploadDokumenPendukung)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

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

// func adminMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		//Split by space " "
// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		//validate token
// 		token, err := authService.ValidateToken(tokenString)

// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		userID := int(claim["user_id"].(float64))

// 		user, err := userService.GetUserByID(userID)

// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		if user.LoginAs != 2 { //must be officer/staff
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		//context currentUser
// 		c.Set("currentUser", user)

// 	}

// }
