package service

import (
    "github.com/go-redis/redis"
	"douyin/database/entity"
	"douyin/database/mysql"
)

// mysql
func SetFollow(fromuid uint64, touid uint64, action uint64) (bool, error) {
	exist, err := FindFollow(fromuid, touid)
	if err != nil {
		return false, err
	}
	db := mysql.Conn()
	tx := db.Begin()
	if action == 2 && exist == true {
		if err := db.Model(&entity.RUserFollow{}).Where(&entity.RUserFollow{Fromuid: fromuid, Touid: touid}).Update("Isdel", 1).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	} else if action == 1 && exist == false {
		rUserFollow := &entity.RUserFollow{
			Fromuid: fromuid,
			Touid:   touid,
		}
		if err := db.Create(rUserFollow).Error; err != nil {
			tx.Rollback()
			return false, nil
		}
	}
	tx.Commit()
	return true, nil
}

func FindFollow(fromuid uint64, touid uint64) (bool, error) {
	var count int64
	db := mysql.Conn()
	tx := db.Begin()
	err := db.Model(&entity.RUserFollow{}).Where(&entity.RUserFollow{Fromuid: fromuid, Touid: touid}).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()

	if count == 0 {
		return false, nil
	}
	return true, nil
}

// follow userid
func GetFavouriteFollower(touid uint64) (map[string]interface{}, error) {
	var res []entity.RUserFollow
	db := mysql.Conn()
	tx := db.Begin()
	err := db.Model(&entity.RUserFollow{}).Where(&entity.RUserFollow{Touid: touid}).Find(&res).Error
	if err != nil {
		tx.Rollback()
		favouriteFollower := map[string]interface{}{"user_list": res}
		return favouriteFollower, err
	}
	tx.Commit()
	favouriteFollower := map[string]interface{}{"user_list": res}
	return favouriteFollower, err
}

// get the user i like(userid follow)
func GetFavouriteFollow(fromuid uint64) (map[string]interface{}, error) {
	var res []entity.RUserFollow
	db := mysql.Conn()
	tx := db.Begin()
	err := db.Model(&entity.RUserFollow{}).Where("fromuid = ? and isdel = ?", fromuid, 0).Find(&res).Error
	if err != nil {
		tx.Rollback()
		favouriteFollow := map[string]interface{}{"user_list": res}
		return favouriteFollow, err
	}
	tx.Commit()
	favouriteFollow := map[string]interface{}{"user_list": res}
	return favouriteFollow, nil
}

func SetVedio(uid uint64, vid uint64, action uint64) (bool, error) {
    exist, err := FindVedio(uid, vid)
	if err != nil {
		return false, err
	}

	db := mysql.Conn()
	tx := db.Begin()
	if action == 1 && exist == false {
		rVideoFavorite := &entity.RVideoFavorite{
			Uid:   uid,
			Vid:   vid,
			Isdel: 0,
		}
		if err := db.Create(rVideoFavorite).Error; err != nil {
			tx.Rollback()
			return false, nil
		}
	} else if action == 2 {
		if exist == true {
			if err := db.Model(&entity.RVideoFavorite{}).Where(&entity.RVideoFavorite{Uid: uid, Vid: vid}).Update("Isdel", 1).Error; err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}
	tx.Commit()
	return true, nil
}

func FindVedio(uid uint64, vid uint64) (bool, error) {
	var count int64
	db := mysql.Conn()
	tx := db.Begin()
	err := db.Model(&entity.RVideoFavorite{}).Where(&entity.RVideoFavorite{Uid: uid, Vid: vid}).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// 找到喜欢我的视频的数目
func CountVedio(vid uint64) (int64, error) {
	var count int64
	db := mysql.Conn()
	tx := db.Begin()
	err := db.Model(&entity.RUserVideo{}).Where(&entity.RUserVideo{Vid: vid}).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return count, err
}

//  favourite vedio list
func GetFavouriteVedio(uid uint64) (map[string]interface{}, error) {
	var res []entity.RVideoFavorite
	db := mysql.Conn()
	tx := db.Begin()
    err := db.Where("uid = ? and isdel = ?", uid, 0).Find(&res).Error
	if err != nil {
		tx.Rollback()
		favouriteList := map[string]interface{}{"video_list": res}
		return favouriteList, err
	}
	tx.Commit()
	favouriteList := map[string]interface{}{"video_list": res}
	return favouriteList, err
}

//redis***************************
func createClient() *redis.Client {

	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// check whether can connect the resid
	_, err := client.Ping().Result()
	if err != nil {

		panic(err)
	}

	return client
}

// find whether the userID in the tableName
func is_like_id(client *redis.Client, tableName string, userID int) bool {
	val, err := client.SIsMember(tableName, userID).Result()

	if err != nil {
		panic(err)
	}

	return val
}

func like_id_set(client *redis.Client, tableName string, userID int) {
	// 点赞和取消点赞
	val := is_like_id(client, tableName, userID)

	if val == false {
		_, err := client.SAdd(tableName, userID).Result()
		if err != nil {
			panic(err)
		}
	} else {
		_, err := client.SRem(tableName, userID).Result()
		if err != nil {
			panic(err)
		}
	}
}

// count
func like_id_count(client *redis.Client, tableName string) int64 {
	val, err := client.SCard(tableName).Result()

	if err != nil {
		panic(err)
	}

	return val
}
