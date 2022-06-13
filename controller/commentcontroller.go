package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "douyin/util"
    "douyin/service"
)

//{"url":"/douyin/comment/list/","function":"1.get videos' all comments, 2.to sort by latesttime reverse order, 3.single maximum use 15"}
func GetVideoCommentList(ctx *gin.Context) {
    var errmsg string
    videoid, err := strconv.ParseUint(ctx.Query("video_id"), 10, 64)
    if err != nil {
        errmsg = "no video"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    if videoid < 1 {
        errmsg = "wrong video"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }

    data, errmsg := service.GetVideoCommentList(videoid)
    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}

//{"url":"/douyin/comment/action/","function":"add videos' comment"}
func AddVideoComment(ctx *gin.Context) {
    var errmsg string
    var data map[string]interface{}
    videoid, err := strconv.ParseUint(ctx.Query("video_id"), 10, 64)
    if err != nil {
        errmsg = "no video"
    }
    actiontype, err := strconv.Atoi(ctx.Query("action_type"))
    if err != nil {
        errmsg = "no type"
    }
    userid := ctx.GetUint64("userid")
    if actiontype == ADDCOMMENTTYPE {
        content := ctx.Query("comment_text")
        if content == "" {
            errmsg = "content empty"
        } else {
            data, errmsg = service.AddVideoComment(userid, videoid, content)
        }
    } else if actiontype == DELCOMMENTTYPE {
        commentid, err := strconv.ParseUint(ctx.Query("comment_id"), 10, 64)
        if err != nil {
            errmsg = "no commentid"
        } else {
            data, errmsg = service.DelVideoComment(userid, videoid, commentid)
        }
    } else {
        errmsg = "wrong type"
    }

    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}

//{"url":"/douyin/reply/list/","function":"1.get comments' all reply, 2.to sort by latesttime reverse order, 3.single maximum use 15"}
func GetCommentReplyList(ctx *gin.Context) {
    var errmsg string
    commentid, err := strconv.ParseUint(ctx.Query("comment_id"), 10, 64)
    if err != nil {
        errmsg = "no comment"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    if commentid < 1 {
        errmsg = "wrong comment"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }

    data, errmsg := service.GetCommentReplyList(commentid)
    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}

//{"url":"/douyin/reply/action/","function":"add comments' reply"}
func AddCommentReply(ctx *gin.Context) {
    var errmsg string
    var data map[string]interface{}
    commentid, err := strconv.ParseUint(ctx.Query("comment_id"), 10, 64)
    if err != nil {
        errmsg = "no comment"
    }
    toid, err := strconv.ParseUint(ctx.Query("to_id"), 10, 64)
    if err != nil {
        errmsg = "no to comment"
    }
    actiontype, err := strconv.Atoi(ctx.Query("action_type"))
    if err != nil {
        errmsg = "no type"
    }
    tempidtype, err := strconv.ParseUint(ctx.Query("idtype"), 10, 8)
    if err != nil {
        errmsg = "no reply"
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, nil))
        return
    }
    idtype := uint8(tempidtype)

    userid := ctx.GetUint64("userid")
    if actiontype == ADDCOMMENTTYPE {
        content := ctx.Query("reply_text")
        if content == "" {
            errmsg = "content empty"
        } else {
            data, errmsg = service.AddCommentReply(userid, commentid, toid, idtype, content)
        }
    } else if actiontype == DELCOMMENTTYPE {
        fromid, err := strconv.ParseUint(ctx.Query("from_id"), 10, 64)
        if err != nil {
            errmsg = "no commentid"
        } else {
            data, errmsg = service.DelCommentReply(fromid)
        }
    } else {
        errmsg = "wrong type"
    }

    if errmsg == "" {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_SUCCESS, errmsg, data))
        return
    } else {
        ctx.JSON(http.StatusOK, util.ResponseJSON(STATUS_FAIL, errmsg, data))
        return
    }
}
