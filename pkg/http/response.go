package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, msg string, data ...interface{}) {
	if len(data) < 1 {
		data = []any{nil}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": msg,
		"data":    data[0],
		"time":    time.Now().Unix(),
	})
}

func Error(c *gin.Context, msg string, data ...interface{}) {
	if len(data) < 1 {
		data = []any{nil}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    500,
		"message": msg,
		"data":    data[0],
		"time":    time.Now().Unix(),
	})
}
