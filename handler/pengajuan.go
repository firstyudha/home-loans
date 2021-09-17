package handler

import (
	"home-loans/helper"
	"home-loans/pengajuan"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pengajuanHandler struct {
	service pengajuan.Service
}

func NewPengajuanHandler(service pengajuan.Service) *pengajuanHandler {
	return &pengajuanHandler{service}
}

func (h *pengajuanHandler) GetPengajuans(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	pengajuans, err := h.service.GetPengajuans(userID)
	if err != nil {
		response := helper.APIResponse("Error to get pengajuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of pengajuans", http.StatusOK, "success", pengajuan.FormatPengajuans(pengajuans))
	c.JSON(http.StatusOK, response)
}
