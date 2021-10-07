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
	userID, _ := strconv.Atoi(c.Query("user_id"))
	kelengkapans, err := h.kelengkapanService.GetKelengkapans(userID)
	if err != nil {
		response := helper.APIResponse("Error to get kelengkapans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of kelengkapans", http.StatusOK, "success", kelengkapan.FormatKelengkapans(kelengkapans))
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

func (h *kelengkapanHandler) DeleteKelengkapan(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	err := h.kelengkapanService.DeleteKelengkapan(userID)
	if err != nil {
		data := gin.H{"is_deleted": false}
		response := helper.APIResponse("Failed to delete kelengkapan", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_deleted": true}
	response := helper.APIResponse("Kelengkapan deleted", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
