package util

import (
    "github.com/gin-gonic/gin"
)

func ResponseJSON(status_code int32, status_msg string, data map[string]interface{}) gin.H {
    returndata := gin.H{
        "status_code": status_code,
        "status_msg": status_msg,
    }

    for k,v := range data {
        returndata[k] = v
    }

    return returndata
}
