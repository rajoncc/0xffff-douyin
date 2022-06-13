package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "douyin/util"
    "douyin/service"
)

func IndexPage(ctx *gin.Context) {
    ctx.String(http.StatusOK, `have:
    GET: /douyin/feed/
         /douyin/user/
         /douyin/publish/list/
         /douyin/favorite/list/
         /douyin/comment/list/
         /douyin/relation/follow/list/
         /douyin/relation/follower/list/
         /douyin/reply/list/

    POST: /douyin/user/register/
          /douyin/user/login/
          /douyin/publish/action/
          /douyin/favorite/action/
          /douyin/comment/action/
          /douyin/relation/action/
          /douyin/reply/action/
    `)
}

//{"url":"/douyin/user/register/","function":"user register"}
func Register(ctx *gin.Context) {
    username := ctx.Query("username")
    password := ctx.Query("password")

    //params validate
    ok, errmsg := util.ValidateUsernamePassword(username, password)
    if !ok {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }

    data, errmsg := service.Register(username, password)
    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}

//{"url":"/douyin/user/login/","function":"user login"}
func Login(ctx *gin.Context) {
    username := ctx.Query("username")
    password := ctx.Query("password")

    //params validate
    ok, errmsg := util.ValidateUsernamePassword(username, password)
    if !ok {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    
    data, errmsg := service.Login(username, password)
    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}

//{"url":"/douyin/user/","function":"1.get user's basic infomation, 2.get user's follow count and follwer count"}
func GetUserInfo(ctx *gin.Context) {
    user_id := ctx.Query("user_id")
    errmsg := ""
    if user_id == "" {
        errmsg = "empty id"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    touserid, err := strconv.ParseUint(user_id, 10, 64)
    if err != nil {
        errmsg = "wrong id format"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    userid := ctx.GetUint64("userid")
    if userid < 1 || touserid < 1 {
        errmsg = "wrong id"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }

    data, errmsg := service.GetUserInfo(userid, touserid)
    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}
