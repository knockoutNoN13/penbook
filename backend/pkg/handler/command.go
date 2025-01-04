package handler

import (
	"net/http"
	"pentbook/pkg/models"

	"github.com/gin-gonic/gin"
)

type getAllCommandsResponse struct {
	Data []models.GetAllResponse `json:"data"`
}

func (h *Handler) getAllCommands(c *gin.Context) {

	commands, err := h.services.Command.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "No commands found")
		return
	}

	response := getAllCommandsResponse{Data: commands}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) getSelectedCommand(c *gin.Context) {

	commandId := c.Param("id")

	command, err := h.services.Command.GetById(commandId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, command)

}

func (h *Handler) createCommand(c *gin.Context) {
	var command models.Command

	err := c.BindJSON(&command)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad json")
		return
	}

	commandId, err := h.services.Command.Create(command)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to create command")
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success",
		"commandId": commandId})
}

func (h *Handler) deleteCommand(c *gin.Context) {

	commandId := c.Param("id")

	err := h.services.Command.Delete(commandId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
