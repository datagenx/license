package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/datagenx/license-generator/internal/generate"
	"github.com/datagenx/license-generator/internal/storage"
)

func PostGenerate(ctx *gin.Context) {

	// reading and binding the post data
	inpBody := generate.Rlic{}

	// Validating the input data
	if err := ctx.ShouldBindJSON(&inpBody); err != nil {
		errMsg(ctx, http.StatusNoContent, err)
		return
	}

	if err := inpBody.InputValidation(); err != nil {
		errMsg(ctx, http.StatusPartialContent, err)
		return
	}

	//generating the license
	sl, lic, err := inpBody.Generate()

	if err != nil {
		errMsg(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	// Writing to Plugins
	if err = storage.Plugins(sl, lic); err != nil {
		errMsg(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	msg(ctx, http.StatusOK, lic)
}

func msg(ctx *gin.Context, status int, message string) {
	ctx.String(status, message)
}

func errMsg(ctx *gin.Context, status int, err error) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"error": err.Error(),
		"code":  status,
	})
}
