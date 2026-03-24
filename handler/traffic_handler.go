package handler

import (
	"backend-noted/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrafficHandler struct {
	Service domain.TrafficService
}

func NewTrafficHandler(r *gin.Engine, service domain.TrafficService) {
	handler := &TrafficHandler{Service: service}
	r.GET("/stats", handler.GetServerStats)
}

func (h *TrafficHandler) GetServerStats(c *gin.Context) {
	stats, err := h.Service.GetServerStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "gagal",
			"error":  "Gagal mengambil data statistik server",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"data":   stats,
	})
}