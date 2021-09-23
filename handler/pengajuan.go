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

	//status default 1 = diajukkan
	input.Status = "1"

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

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
	var inputID pengajuan.GetPengajuanInput

	err := c.ShouldBindUri(&inputID)

	var inputData pengajuan.CreatePengajuanInput

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload bukti ktp", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
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

	_, err = h.pengajuanService.SaveBuktiKTP(inputID, path)
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
	var inputID pengajuan.GetPengajuanInput

	err := c.ShouldBindUri(&inputID)

	var inputData pengajuan.CreatePengajuanInput

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload bukti slip gaji", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
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

	_, err = h.pengajuanService.SaveBuktiSlipGaji(inputID, path)
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
