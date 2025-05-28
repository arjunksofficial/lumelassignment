package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) TriggerImport(c *gin.Context) {
	h.logger.Info("Triggering data import...")
	c.JSON(200, gin.H{
		"message": "Data import triggered successfully",
	})
}
