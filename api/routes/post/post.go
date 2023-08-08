package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rowing-market-api/api/models"
	"rowing-market-api/api/services"
	"rowing-market-api/api/validations"
	"rowing-market-api/pkg/apierror"
	headers "rowing-market-api/pkg/header"
	"rowing-market-api/pkg/logger"
)

func RegisterPost(c *gin.Context) {
	lang := headers.GetAcceptLanguage(c)
	var param models.Post

	err := c.BindJSON(&param)
	ctx := c.Request.Context()
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("RegisterClub BindJSON: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, apierror.CreateError(apierror.CodeValidationFailed, apierror.MsgValidationFailed, "en"))
		return
	}

	errMsg := validations.CreatePostValidation(ctx, param, lang)
	if errMsg.HasError {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": errMsg})
		return
	}

	result, err := services.GetPostService(ctx, nil).SavePost(param)
	created := false
	if result == 1 {
		created = true
	}
	c.JSON(http.StatusOK, gin.H{"created": created})
}
