package controller

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"

    "douyin/util"
    "douyin/service"
)

//{"url":"/douyin/feed/","function":"1.get feedflow, 2.to sort by reverse order latesttime , 3.single maximum 30, we can use 15"}
func GetIndexFeedFlow(ctx *gin.Context) {
    // latest_time
	var latest_time int64
	latest_time_str := ctx.Query("latest_time")
	if latest_time_str != "" {
		var err error
		latest_time, err = strconv.ParseInt(latest_time_str, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "lastest_time paese error: "+err.Error(), nil))
			return
		}
	} else {
		latest_time = int64(time.Now().Unix())
	}

	// token
	var vlist []map[string]interface{}
	var next_time int64
	var err error
	token := ctx.Query("token")
	if token != "" {
		user_id := ctx.GetUint64("userid")
		vlist, next_time, err = service.GetIndexFeedFlow(GET_VIDEO_NUMBER, latest_time, user_id, true)
	} else {
		vlist, next_time, err = service.GetIndexFeedFlow(GET_VIDEO_NUMBER, latest_time, uint64(0), false)
	}

	if err != nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "get feedlist success",
		map[string]interface{}{
			"next_time":  next_time,
			"video_list": vlist,
		})) 
}

//{"url":"/douyin/publish/list/","function":"get all feedflow from current user"}
func GetUserFeedList(ctx *gin.Context) {
    token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty token", nil))
		return
	}

	user_id_str := ctx.Query("user_id")
	if user_id_str == "" {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty user_id", nil))
		return
	}
	user_id, err := strconv.ParseUint(user_id_str, 10, 64) // 要获取该用户的视频
	if err != nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "user_id format error", nil))
		return
	}

	login_user_id := ctx.GetUint64("userid") // 当前登陆用户

	vlist, err := service.GetUserFeedList(login_user_id, user_id)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "get user feedlist success",
		map[string]interface{}{"video_list": vlist}))
}

//{"url":"/douyin/publish/action/","function":"1.upload video, 2.add to user's feedflow"}
func AddVideoFeedFlow(ctx *gin.Context) {
    title := string(ctx.PostForm("title"))
	if title == "" {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty title", nil))
		return
	}

	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty data", nil))
		return
	}

	token := string(ctx.PostForm("token"))
	if token == "" {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, "empty token", nil))
		return
	}

	err = service.AddVideoFeedFlow(ctx, file, UPLOAD_PATH, token, title)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, err.Error(), nil))
		return
	} else {
		ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, "add feedflow success", nil))
		return
	}
}
