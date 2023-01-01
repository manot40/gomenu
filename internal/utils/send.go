package utils

import (
	"github.com/gin-gonic/gin"
)

func SendJson(ctx *gin.Context, status int, data interface{}, message ...string) {
	resMap := map[string]interface{}{"code": status, "message": StatusMessage[status]}

	if message != nil {
		resMap["message"] = message[0]
	}

	if status >= 400 {
		resMap["error"] = data
	} else {
		resMap["result"] = data
	}

	meta, exist := ctx.Get("meta")
	if exist {
		resMap["meta"] = meta
	}

	ctx.JSON(status, resMap)
}

var StatusMessage = map[int]string{
	200: "Request Processed",
	201: "Resource Created/Updated",
	202: "Accepted",
	204: "No Content Found",
	400: "Bad request, please check your submit entry",
	401: "You are unauthorized to access this resource",
	403: "Access to this resource forbidden",
	404: "Requested Resource Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	409: "Conflict",
	410: "Gone",
	422: "Unprocessable Entity",
	500: "Server can't processing your request at this moment",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Server can't processing your request at this moment",
	504: "Gateway Timeout",
}
