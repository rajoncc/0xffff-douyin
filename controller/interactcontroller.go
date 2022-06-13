package controller

import (
    "douyin/service"
	"douyin/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//{"url":"/douyin/favorite/list/","function":"get user's all favorite feedflow"}
func GetUserFavoriteList(ctx *gin.Context) {
    user_id := ctx.Query("user_id")
	errmsg := ""
	if user_id == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	uid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	data, err := service.GetFavouriteVedio(uid)

	if err == nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, " success getuserfavoritelist test!", data))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "fail getuserfavoritelist test!", data))
		return
	}
}

//{"url":"/douyin/favorite/action/","function":"1.favorite action, 2.cancel favorite action"}
func FavoriteAction(ctx *gin.Context) {
    vedio_id := ctx.Query("video_id")
	errmsg := ""
	if vedio_id == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	vid, err := strconv.ParseUint(vedio_id, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	action := ctx.Query("action_type")
	if action == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty id", nil))
		return
	}
	action_type, err := strconv.ParseUint(action, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	uid := ctx.GetUint64("userid")
	if vid < 1 || uid < 1 {
		errmsg = "wrong id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	data, err := service.SetVedio(uid, vid, action_type)
	if err == nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "success favoriteaction test!", map[string]interface{}{"success": data}))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, err.Error(), map[string]interface{}{"success": data}))
		return
	}
}

//{"url":"/douyin/relation/follow/list/","function":"get user's follow userlist"}
func GetFollowUserList(ctx *gin.Context) {
user_id := ctx.Query("user_id")
	errmsg := ""
	if user_id == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	uid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	if uid < 1 {
		errmsg = "wrong id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	data, err := service.GetFavouriteFollow(uid)
	if err == nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "success getfollowuserlist test!", data))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "fail getfollowuserlist test!", data))
		return
	}
}

//{"url":"/douyin/relation/follower/list/","function":"get user's follower userlist"}
func GetFollowerUserList(ctx *gin.Context) {
    user_id := ctx.Query("user_id")
	errmsg := ""
	if user_id == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	uid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	if uid < 1 {
		errmsg = "wrong id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	data, err := service.GetFavouriteFollower(uid)
	if err == nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "success getfolloweruserlist test!", data))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "fail getfolloweruserlist test!", data))
		return
	}
}

//{"url":"/douyin/relation/action/","function":"1.follow action, 2.cancel follow action"}
func FollowAction(ctx *gin.Context) {
    user_id := ctx.Query("to_user_id")
	errmsg := ""
	if user_id == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}
	touid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	action := ctx.Query("action_type")
	if action == "" {
		errmsg = "empty id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty id", nil))
		return
	}
	action_type, err := strconv.ParseUint(action, 10, 64)
	if err != nil {
		errmsg = "wrong id format"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	fromuid := ctx.GetUint64("userid")
	if touid < 1 || fromuid < 1 {
		errmsg = "wrong id"
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
		return
	}

	data, err := service.SetFollow(fromuid, touid, action_type)
	if err == nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "success followaction test!", map[string]interface{}{"success": data}))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "fail followaction test!", map[string]interface{}{"success": data}))
		return
	}
}
