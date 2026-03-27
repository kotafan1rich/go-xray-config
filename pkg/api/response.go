package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func SuccessWithMeta(c *gin.Context, data any, meta any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   message,
	})
}

func ValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	})
}

func NotFound(c *gin.Context, resource string) {
	c.JSON(http.StatusNotFound, Response{
		Error: resource + " not found",
	})
}

func BadRequest(c *gin.Context, message string) {
    c.JSON(http.StatusBadRequest, Response{
        Success: false,
        Error:   message,
    })
}

func InternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
        Success: false,
        Error:   message,
    })
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
        Success: false,
        Error:   "Unauthorized",
    })
}

func Forbidden(c *gin.Context) {
    c.JSON(http.StatusForbidden, Response{
        Success: false,
        Error:   "Forbidden",
    })
}