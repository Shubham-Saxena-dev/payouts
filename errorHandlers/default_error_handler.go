package errorHandlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ErrorHandlers interface {
	FailOnError(err error, msg string)
	HandleError(context *gin.Context, err error, statusCode int)
}

type errorHandler struct {
}

func NewErrorHandler() ErrorHandlers {
	return &errorHandler{}
}

func (handler *errorHandler) FailOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}

func (handler *errorHandler) HandleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
