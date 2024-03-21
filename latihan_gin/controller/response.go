package controller

import (
	m "latihan_gin/model"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, kode int, message string) {
	var response m.SuccessResponse
	response.Status = kode
	response.Message = message
}

func SendErrorResponse(c *gin.Context, kode int, message string) {
	var response m.ErrorResponse
	response.Status = kode
	response.Message = message

}
