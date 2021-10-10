package handler

import (
	"fmt"
	"home-loans/auth"
	"home-loans/helper"
	"home-loans/kelengkapan"
	"home-loans/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type kelengkapanHandler struct {
	kelengkapanService kelengkapan.Service
	authService        auth.Service
}

func NewKelengkapanHandler(kelengkapanService kelengkapan.Service, authService auth.Service) *kelengkapanHandler {
	return &kelengkapanHandler{kelengkapanService, authService}
}

func (h *kelengkapanHandler) GetKelengkapans(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	//user not login
	if currentUser.ID == 0 {
		response := helper.APIResponse("Unauthorized.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//login as customer
	if currentUser.LoginAs == 1 {
		kelengkapans, err := h.kelengkapanService.GetKelengkapan(currentUser.ID)
		if err != nil {
			response := helper.APIResponse("Error to get kelengkapans", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Data kelengkapan", http.StatusOK, "success", kelengkapan.FormatKelengkapans(kelengkapans))
		c.JSON(http.StatusOK, response)
	}

	//login as staff
	if currentUser.LoginAs == 2 {
		kelengkapans, err := h.kelengkapanService.GetKelengkapans()
		if err != nil {
			response := helper.APIResponse("Error to get kelengkapans", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("List of kelengkapans", http.StatusOK, "success", kelengkapan.FormatKelengkapans(kelengkapans))
		c.JSON(http.StatusOK, response)
	}

}

func (h *kelengkapanHandler) GetKelengkapanDetail(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	kelengkapans, err := h.kelengkapanService.GetKelengkapan(userID)
	if err != nil {
		response := helper.APIResponse("Error to get kelengkapan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kelengkapan detail", http.StatusOK, "success", kelengkapan.FormatKelengkapans(kelengkapans))
	c.JSON(http.StatusOK, response)
}

func (h *kelengkapanHandler) CreateKelengkapan(c *gin.Context) {
	var input kelengkapan.CreateKelengkapanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create kelengkapan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//status default 1 = diajukkan
	input.Status = "1"

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newKelengkapan, err := h.kelengkapanService.CreateKelengkapan(currentUser.ID, input)
	if err != nil {
		data := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to create kelengkapan", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create kelengkapan", http.StatusOK, "success", kelengkapan.FormatKelengkapan(newKelengkapan))
	c.JSON(http.StatusOK, response)
}

func (h *kelengkapanHandler) UploadDokumenPendukung(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	file, err := c.FormFile("dokumen_pendukung")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload dokumen pendukung", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file_name := strings.ReplaceAll(file.Filename, " ", "-") //remove whitespace

	path := fmt.Sprintf("dokumen_pendukung_files/%d-%s", userID, file_name)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload dokumen pendukung", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.kelengkapanService.SaveDokumenPendukung(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload dokumen pendukung", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Dokumen pendukung successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *kelengkapanHandler) UpdateKelengkapanStatus(c *gin.Context) {
	var userID kelengkapan.GetKelengkapanInput

	err := c.ShouldBindUri(&userID)
	if err != nil {
		response := helper.APIResponse("Failed to update kelengkapan status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData kelengkapan.UpdateKelengkapanStatusInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update kelengkapan status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedKelengkapanStatus, err := h.kelengkapanService.SaveKelengkapanStatus(userID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update kelengkapan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update kelengkapan", http.StatusOK, "success", kelengkapan.FormatKelengkapan(updatedKelengkapanStatus))
	c.JSON(http.StatusOK, response)

}
