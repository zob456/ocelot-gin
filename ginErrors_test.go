package ocelotGin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)


func testGinAutErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		OErrorHandler(ctx, NotAuthorizedErr, NotAuthorizedCode)
	}
}

func testGinBadRequestErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		OErrorHandler(ctx, BadRequestErr, BadRequestCode)
	}
}

func testGinAuthSqlErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		OAuthSqlErrorHandler(ctx, TestSqlErr)
	}
}

func testGinSqlErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		OSqlErrorHandler(ctx, TestSqlErr)
	}
}

func testGinExpectedSqlErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		OExpectedNoRowsInSqlErrorHandler(ctx, TestSqlErr)
	}
}


func Test_OcelotGinErrorHandler_NotAuthorized_HAPPY(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinAutErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.Equal(t, NotAuthorizedCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Not_Authorized_SAD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinAutErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.NotEqual(t, InternalServerErrCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Bad_Reuqest_HAPPY(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinBadRequestErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.Equal(t, BadRequestCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Bad_Request_SAD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinBadRequestErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.NotEqual(t, InternalServerErrCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Sql_HAPPY(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.Equal(t, NotFoundCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Sql_SAD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.NotEqual(t, InternalServerErrCode, rr.Code)
}

func Test_OGinErrorHandler_Auth_Sql_HAPPY(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinAuthSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.Equal(t, NotAuthorizedCode, rr.Code)
}

func Test_OGinErrorHandler_Auth_Sql_SAD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinAuthSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.NotEqual(t, InternalServerErrCode, rr.Code)
}

func Test_OcelotGinErrorHandler_Expected_Sql_HAPPY(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinExpectedSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
}

func Test_OcelotGinErrorHandler_Expected_Sql_SAD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	rr := httptest.NewRecorder()
	r.POST("/", testGinExpectedSqlErrorHandler())
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(rr, request)
	assert.NotEqual(t, InternalServerErrCode, rr.Code)
}