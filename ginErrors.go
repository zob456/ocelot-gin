package ocelotGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zob456/ocelot-logger"
	"log"
)


func OErrorHandler(ctx *gin.Context, err error, errorCode int ) {
	publicErrMessage := ReturnGinPublicErrorMessage(errorCode)
	ocelotLogger.ErrorLogger(err)
	ctx.AbortWithStatusJSON(errorCode, publicErrMessage)
	return
}

func OAuthSqlErrorHandler(ctx *gin.Context, err error) {
	if err.Error() == SqlErr {
		ocelotLogger.ErrorLogger(err)
		ctx.AbortWithStatusJSON(401, PublicNotAuthorized)
		return
	}
	log.Println(fmt.Sprintf("ERROR: %+v", err))
	ctx.AbortWithStatusJSON(500, PublicInternalServerError)
	return
}

func OSqlErrorHandler(ctx *gin.Context, err error) {
		if err.Error() == SqlErr {
			ocelotLogger.ErrorLogger(err)
			ctx.AbortWithStatusJSON(404, PublicNotFound)
			return
		}
		log.Println(fmt.Sprintf("ERROR: %+v", err))
		ctx.AbortWithStatusJSON(500, PublicInternalServerError)
		return
}

func OExpectedNoRowsInSqlErrorHandler(ctx *gin.Context, err error) {
	if err.Error() != SqlErr {
		ocelotLogger.ErrorLogger(err)
		ctx.AbortWithStatusJSON(500, PublicInternalServerError)
		return
	}
}