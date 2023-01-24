package handler

import (
	"github.com/gin-gonic/gin"
	"microServices/bet_maker"
	"net/http"
)

func (h *Handler) createEvent(c *gin.Context) {
	var input bet_maker.Event

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateEvent(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
