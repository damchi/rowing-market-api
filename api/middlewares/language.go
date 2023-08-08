package middlewares

import (
	"github.com/gin-gonic/gin"
	headers "rowing-market-api/pkg/header"
)

var GlobalLang string

func Language() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := headers.GetAcceptLanguage(c)
		GlobalLang = lang

		c.Next()
	}
}
