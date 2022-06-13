package middleware

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    "douyin/database/mysql"
    "douyin/database/entity"
    "douyin/util"
)

func TokenCheck(ctx *gin.Context) {
    token, ok := ctx.GetQuery("token")
    if !ok || token == "" {
        errmsg := "no token"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL_REDIRECT, errmsg, nil))
        ctx.Abort()
        return
    }

    db := mysql.Conn()
    var usertoken entity.UserToken
    ret := db.Select("uid", "token", "expiredtime").Take(&usertoken, "token = ?", token)
    if ret.Error != nil {
        errmsg := "not right token"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL_REDIRECT, errmsg, nil))
        ctx.Abort()
        return
    }
    
    if usertoken.Expiredtime > time.Now().Unix() {
        ctx.Set("userid", usertoken.Uid)//save uid from token,may be used to compare with userid to validata is the same user.
    } else {
        errmsg := "old token"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL_REDIRECT, errmsg, nil))
        ctx.Abort()
        return
    }
}
