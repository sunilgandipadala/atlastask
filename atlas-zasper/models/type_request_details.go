package models

import "github.com/gin-gonic/gin"

type RequestDetails struct {
	URL      string
	Method   string
	Request  *gin.Context
	Function func(*gin.Context)
}
