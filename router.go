package main

import (
    "github.com/gin-gonic/gin"

    "douyin/middleware"
    "douyin/controller"
)

func setRouter(router *gin.Engine) {
    router.Static("/feedfile", "./static/feedfile")

    douyin := router.Group("/douyin")
    {
        notoken := douyin.Group("")
        {
            notoken.GET("/", controller.IndexPage)
            //basic apis
            notoken.GET("/feed/", controller.GetIndexFeedFlow)
            notoken.POST("/user/register/", controller.Register)
            notoken.POST("/user/login/", controller.Login)
        }
        needtoken := douyin.Group("", middleware.TokenCheck)
        {
            //basic apis
            needtoken.GET("/user/", controller.GetUserInfo)
            needtoken.GET("/publish/list/", controller.GetUserFeedList)
            needtoken.POST("/publish/action/", controller.AddVideoFeedFlow)

            //extra apis - I
            needtoken.GET("/favorite/list/", controller.GetUserFavoriteList)
            needtoken.GET("/comment/list/", controller.GetVideoCommentList)
            needtoken.POST("/favorite/action/", controller.FavoriteAction)
            needtoken.POST("/comment/action/", controller.AddVideoComment)

            //extra apis - II
            needtoken.GET("/relation/follow/list/", controller.GetFollowUserList)
            needtoken.GET("/relation/follower/list/", controller.GetFollowerUserList)
            needtoken.POST("/relation/action/", controller.FollowAction)

            //extra apis - III
            needtoken.GET("/reply/list/", controller.GetCommentReplyList)  //comments' reply
            needtoken.POST("/reply/action/", controller.AddCommentReply)    //add comments' reply
        }
    }
}
