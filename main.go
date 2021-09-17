package main

import (
	"home-loans/auth"
	"home-loans/handler"
	"home-loans/kelengkapan"
	"home-loans/pengajuan"
	"home-loans/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3309)/golang-homeloans?charset=utf8mb4&parseTime=True&loc=Local"
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
	pengajuanHandler := handler.NewPengajuanHandler(pengajuanService)

	router := gin.Default()

	//publish file
	router.Static("pengajuan_files/", "./pengajuan_files")
	router.Static("kelengkapan_files/", "./kelengkapan_files")
	router.Static("bukti_ktp_files/", "./bukti_ktp_files")
	router.Static("bukti_slip_gaji_files/", "./bukti_slip_gaji_files")
	router.Static("dokumen_pendukung_files/", "./dokumen_pendukung_files")

	//grouping
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	//pengajuan endpoint
	api.GET("/pengajuan", pengajuanHandler.GetPengajuans)

	router.Run()

}

// func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

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

// 		//context currentUser
// 		c.Set("currentUser", user)

// 	}

// }
