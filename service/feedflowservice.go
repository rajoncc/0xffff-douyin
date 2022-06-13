package service

import (
    "douyin/database/entity"
	"douyin/database/mysql"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoFlowInfo struct {
    Id int64 `json:"id"`
    Author UserInfo `json:"author"`
    Playurl string `json:"play_url"`
    Coverurl string `json:"cover_url"`
    Favoritecount uint32 `json:"favorite_count"`
    Commentcount uint32 `json:"comment_count"`
    IsFavorite bool `json:"is_favorite"`
    Title string `json:"title"`
}

func GetIndexFeedFlow(video_number int, latest_time int64, user_id uint64, is_login_user bool) ([]map[string]interface{}, int64, error) {
defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	videoInfoList := make([]map[string]interface{}, 0, video_number)

	db := mysql.Conn()

	Videos := make([]entity.Video, 0, video_number)
	result := db.Where("Createtime < ?", latest_time).Order("Createtime desc").Limit(video_number).Find(&Videos)
	if result.RowsAffected == 0 {
		return nil, latest_time, errors.New("no video")
	}
	next_time := Videos[len(Videos)-1].Createtime

	for _, Video := range Videos {
		videoInfo := make(map[string]interface{})

		// video
		videoInfo["id"] = Video.Vid
		videoInfo["play_url"] = Video.Playurl
		videoInfo["cover_url"] = Video.Coverurl
		videoInfo["title"] = Video.Title

		// videocount
		VideoCount := entity.VideoCount{}
		db.Where("Vid = ?", Video.Vid).First(&VideoCount)
		videoInfo["favorite_count"] = VideoCount.Favoritecount
		videoInfo["comment_count"] = VideoCount.Commentcount

		// is_favorite
		if is_login_user {
			RVideoFavorite := entity.RVideoFavorite{}
			result := db.Where("Uid = ?, Vid = ?", user_id, Video.Vid).First(&RVideoFavorite)
			if result.RowsAffected == 0 {
				videoInfo["is_favorite"] = false
			} else {
				videoInfo["is_favorite"] = true
			}
		}

		// author
		RUserVideo := entity.RUserVideo{}
		db.Where("Vid = ?", Video.Vid).First(&RUserVideo)

		User := entity.User{}
		db.Where("Uid = ?", RUserVideo.Uid).First(&User)

		UserCount := entity.UserCount{}
		db.Where("Uid = ?", RUserVideo.Uid).First(&UserCount)

		author := make(map[string]interface{})
		author["id"] = RUserVideo.Uid
		author["name"] = User.Uname
		author["follow_count"] = UserCount.Followcount
		author["follower_count"] = UserCount.Followercount
		if is_login_user {
			RUserFollow := entity.RUserFollow{}
			result := db.Where("Fromuid = ?, Touid = ?", user_id, RUserVideo.Uid).First(&RUserFollow)
			if result.RowsAffected == 0 {
				author["is_follow"] = false
			} else {
				author["is_follow"] = true
			}
		} else {
			author["is_follow"] = false
		}
		videoInfo["author"] = author
		videoInfoList = append(videoInfoList, videoInfo)
	}

	return videoInfoList, next_time, nil
}

func GetUserFeedList(login_user_id, user_id uint64) ([]map[string]interface{}, error) {
    defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	videoInfoList := make([]map[string]interface{}, 0)

	db := mysql.Conn()
	RUserVideos := make([]entity.RUserVideo, 0)
	result := db.Where("Uid = ?", user_id).Find(&RUserVideos)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("no video")
	}

	for _, RUserVideo := range RUserVideos {
		videoInfo := make(map[string]interface{})

		// video
		Video := entity.Video{}
		db.Where("Vid = ?", RUserVideo.Vid).First(&Video)
		videoInfo["id"] = Video.Vid
		videoInfo["play_url"] = Video.Playurl
		videoInfo["cover_url"] = Video.Coverurl
		videoInfo["title"] = Video.Title

		// videocount
		VideoCount := entity.VideoCount{}
		db.Where("Vid = ?", Video.Vid).First(&VideoCount)
		videoInfo["favorite_count"] = VideoCount.Favoritecount
		videoInfo["comment_count"] = VideoCount.Commentcount

		// is_favorite
		RVideoFavorite := entity.RVideoFavorite{}
		result := db.Where("Uid = ? and Vid = ?", user_id, Video.Vid).First(&RVideoFavorite)
		if result.RowsAffected == 0 {
			videoInfo["is_favorite"] = false
		} else {
			videoInfo["is_favorite"] = true
		}

		// author
		if login_user_id != user_id {
			User := entity.User{}
			db.Where("Uid = ?", user_id).First(&User)

			UserCount := entity.UserCount{}
			db.Where("Uid = ?", user_id).First(&UserCount)

			author := make(map[string]interface{})
			author["id"] = user_id
			author["name"] = User.Uname
			author["follow_count"] = UserCount.Followcount
			author["follower_count"] = UserCount.Followercount

			RUserFollow := entity.RUserFollow{}
			result := db.Where("Fromuid = ?, Touid = ?", login_user_id, user_id).First(&RUserFollow)
			if result.RowsAffected == 0 {
				author["is_follow"] = false
			} else {
				author["is_follow"] = true
			}
			videoInfo["author"] = author
		}

		videoInfoList = append(videoInfoList, videoInfo)
	}
	return videoInfoList, nil
}

func AddVideoFeedFlow(ctx *gin.Context, file *multipart.FileHeader, upload_path string, token string, title string) error {
    db := mysql.Conn()
	tx := db.Begin()
	err := tx.Error
	if err != nil {
		return errors.New("transaction begin wrong")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	UserToken := entity.UserToken{}
	db.Where("Token = ?", token).First(&UserToken)
	user_id := UserToken.Uid

    // 视频在本地的存储路径为，UPLOAD_PATH + user_id_str + "/" + video_id_str
    now := time.Now().Unix()
    user_id_str := strconv.FormatUint(user_id, 10)
    stamp_str := strconv.FormatInt(now, 10)
    filename := user_id_str+"-"+stamp_str+"-"+file.Filename
    if err := ctx.SaveUploadedFile(file, upload_path+"/"+filename); err != nil {
        tx.Rollback()
        return errors.New(err.Error() + ": upload file fail")
    }

	// 1 video
	Video := entity.Video{Title: title, Playurl: filename, Coverurl: filename, Createtime: now}
	result := db.Create(&Video)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 2 videocount
	VideoCount := entity.VideoCount{Vid: Video.Vid, Favoritecount: 0, Commentcount: 0}
	result = db.Create(&VideoCount)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 3 ruservideo table
	RUserVideo := entity.RUserVideo{Uid: user_id, Vid: Video.Vid}
	result = db.Create(&RUserVideo)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

/*	// 视频在本地的存储路径为，UPLOAD_PATH + user_id_str + "/" + video_id_str
	user_id_str := strconv.FormatUint(user_id, 10)
	video_id_str := strconv.FormatUint(Video.Vid, 10)
	if err := ctx.SaveUploadedFile(file, upload_path+"/"+user_id_str+"-"+video_id_str+"-"+file.Filename); err != nil {
		tx.Rollback()
		return errors.New(err.Error() + ": upload file fail")
	}
*/
	tx.Commit()

	return nil
}
