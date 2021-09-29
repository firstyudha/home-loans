package handler

import (
	"fmt"
	"home-loans/auth"
	"home-loans/helper"
	"home-loans/pengajuan"
	"home-loans/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type pengajuanHandler struct {
	pengajuanService pengajuan.Service
	authService      auth.Service
}

func NewPengajuanHandler(pengajuanService pengajuan.Service, authService auth.Service) *pengajuanHandler {
	return &pengajuanHandler{pengajuanService, authService}
}

func (h *pengajuanHandler) GetPengajuans(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	pengajuans, err := h.pengajuanService.GetPengajuans(userID)
	if err != nil {
		response := helper.APIResponse("Error to get pengajuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of pengajuans", http.StatusOK, "success", pengajuan.FormatPengajuans(pengajuans))
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) CreatePengajuan(c *gin.Context) {
	var input pengajuan.CreatePengajuanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	input.UserID = currentUser.ID

	//status default 1 = diajukkan
	input.Status = "1"

	newPengajuan, err := h.pengajuanService.CreatePengajuan(input)
	if err != nil {
		data := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to create pengajuan", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create pengajuan", http.StatusOK, "success", pengajuan.FormatPengajuan(newPengajuan))
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) UploadBuktiKtp(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	file, err := c.FormFile("bukti_ktp")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti ktp", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file_name := strings.ReplaceAll(file.Filename, " ", "-") //remove whitespace

	path := fmt.Sprintf("bukti_ktp_files/%d-%s", userID, file_name)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti ktp", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.pengajuanService.SaveBuktiKTP(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti ktp", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Bukti ktp successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) UploadBuktiSlipGaji(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	file, err := c.FormFile("bukti_slip_gaji")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti slip gaji", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file_name := strings.ReplaceAll(file.Filename, " ", "-") //remove whitespace

	path := fmt.Sprintf("bukti_slip_gaji_files/%d-%s", userID, file_name)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti slip gaji", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.pengajuanService.SaveBuktiSlipGaji(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload bukti slip gaji", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Bukti slip gaji successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) CheckRecommendation(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	recommendation, err := h.pengajuanService.CheckRecommendation(userID)
	if err != nil {
		response := helper.APIResponse("Error to get recommendation", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get recommendation", http.StatusOK, "success", recommendation)
	c.JSON(http.StatusOK, response)
}
