package ocelotGin

import "errors"

const (
	BadRequestCode        = 400
	NotAuthorizedCode     = 401
	NotFoundCode          = 404
	InternalServerErrCode = 500
)

var (
	BadRequestErr    = errors.New(PublicBadRequest)
	NotAuthorizedErr = errors.New(PublicNotAuthorized)
	TestSqlErr = errors.New(SqlErr)
)
