package utils

import "github.com/gin-gonic/gin"

func ErrorMessageResponse(message string) gin.H {
	return gin.H{
		"status":     "failed",
		"statuscode": -1,
		"error": gin.H{
			"message": message,
		},
	}
}

func SuccessMessageResponse(message string) gin.H {
	return gin.H{
		"status":     "success",
		"statuscode": 0,
		"data": gin.H{
			"message": message,
		},
	}
}

// SuccessDataResponse gives a successful request response.
func SuccessDataResponse(data interface{}) gin.H {
	return gin.H{
		"status":     "success",
		"statuscode": 0,
		"data":       data,
	}
}

func SuccessLoginResponse(message string, data interface{}) gin.H {
	return gin.H{
		"status":      "success",
		"statuscode":  0,
		"description": message,
		"data":        data,
	}
}
