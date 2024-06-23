package responses

import "github.com/gin-gonic/gin"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(context *gin.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return context.JSON(statusCode, data)
}

func MessageResponse(c *gin.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
