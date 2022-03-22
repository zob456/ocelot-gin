package ocelotGin

// error codes

const (
	PublicBadRequest string = "Bad request"
	PublicNotFound string = "Not found"
	PublicInternalServerError string = "Internal Server Error"
	PublicNotAuthorized string = "Not authorized"
	SqlErr string = "sql: no rows in result set"
)

func ReturnGinPublicErrorMessage(errCode int) string {
	var err string
	switch errCode {
	case 400:
		err = PublicBadRequest
	case 401:
		err = PublicNotAuthorized
	case 404:
		err = PublicNotFound
	default:
		err = PublicInternalServerError
	}
	return err
}
