package controllers

/*
This is a controller class where all routes are directed to.
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"takeHomeTest/errorHandlers"
	"takeHomeTest/models"
	"takeHomeTest/service"
)

type Controller interface {
	CreatePayout(c *gin.Context)
}

type controller struct {
	service      service.Service
	errorHandler errorHandlers.ErrorHandlers
}

func NewController(service service.Service, errorHandler errorHandlers.ErrorHandlers) Controller {
	return &controller{
		service:      service,
		errorHandler: errorHandler,
	}
}

/*
This method is a rest api controller for create payout POST request.
It the response along with status code
*/

func (c *controller) CreatePayout(context *gin.Context) {

	var item []models.Item
	if err := context.ShouldBindJSON(&item); err != nil {
		c.errorHandler.HandleError(context, err, http.StatusBadRequest)
		return
	}

	response, err := c.service.CreatePayout(item)
	if err != nil {
		c.errorHandler.HandleError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, &models.ApiResponse{
		NoOfTransactions: len(response),
		Payout:           response,
	})

}
