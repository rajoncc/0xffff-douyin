package service

import (
    "fmt"
    "time"
    "strconv"
    "strings"
    "crypto/sha256"

    uuid "github.com/satori/go.uuid"

    "douyin/database/mysql"
    "douyin/database/entity"
)

type UserInfo struct {
    ID uint64 `json:"id"`
    Name string `json:"name"`
    FollowCount uint32 `json:"follow_count"`
    FollowerCount uint32 `json:"follower_count"`
    //IsFollow bool `json:"is_follow"`
    IsFollow int8 `json:"is_follow"`
}

//return {"user_id":1,"token":"test"}
func Register(username string, password string) (map[string]interface{}, string) {
    //user
    now := time.Now().Unix()
    usersalt := uuid.NewV4().String()[0:6] + strconv.FormatInt(now, 10)[0:3]
    realpassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password + usersalt)))
    randstring := strings.Replace(uuid.NewV4().String()[0:16], "-", "", -1)
    nickname := "user" + randstring
    //insert user
    db := mysql.Conn()
    tx := db.Begin()
    err := tx.Error
    if err != nil {
        return nil, "transaction begin wrong"
    }
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()
    //check existed user
    user := entity.User{Uname: username, Pword: realpassword, Salt:usersalt, Nickname: nickname}
    result := tx.Select("uid").Take(&user, "uname = ?", user.Uname)
    if result.Error == nil {
        tx.Rollback()
        return nil, "user existed"
    }

    result = tx.Create(&user)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "user cant register"
    }
    uid := user.Uid
    if uid < 1 {
        tx.Rollback()
        return nil, "user not register"
    }

    //insert usertoken
    token := ""
    var expiredtime int64 = 0
    usertoken := entity.UserToken{Uid: uid, Token: token, Expiredtime: expiredtime}
    result = tx.Create(&usertoken)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "usertoken cant register"
    }

    //insert usercount
    var followcount uint32 = 0
    var followercount uint32 = 0
    usercount := entity.UserCount{Uid: uid, Followcount: followcount, Followercount: followercount}
    result = tx.Create(&usercount)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "userinfo cant register"
    }

    tx.Commit()
    returndata, errmsg := Login(username, password)
    return returndata, errmsg
}

//return {"user_id":1,"token":"test"}
func Login(username string, password string) (map[string]interface{}, string) {
    user := entity.User{}
    now := time.Now()
    randstring := strings.Replace(uuid.NewV4().String(), "-", "", -1)
    stampsha := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.FormatInt(now.Unix(), 10))))
    newtoken := randstring[:23] + stampsha[:18] + randstring[23:]
    expiredtime := now.Add(7 * 24 * time.Hour).Unix()

    //select user
    db := mysql.Conn()
    result := db.Select("uid", "pword", "salt").Take(&user, "uname = ?", username)
    if result.Error != nil {
        return nil, "user not exist"
    }
    //validate password
    password = fmt.Sprintf("%x", sha256.Sum256([]byte(password + user.Salt)))
    if password != user.Pword {
        return nil, "wrong password"
    }

    //update usertoken
    usertoken := entity.UserToken{Uid: user.Uid, Token: newtoken, Expiredtime: expiredtime}
    result = db.Model(&usertoken).Where("uid = ?", user.Uid).Updates(&usertoken)
    if result.Error != nil || result.RowsAffected < 1 {
        return nil, "user token fail"
    }

    returndata := map[string]interface{}{"user_id": user.Uid, "token": newtoken}
    return returndata, ""
}

//return {"user":{"id":1,"name":"test","follow_count":99,"follower_count":"99","is_follow":true}}
func GetUserInfo(userid uint64, touserid uint64) (map[string]interface{}, string) {
    if userid == touserid {
        //userinfo owner or not,different user can see different userinfo
    }

    //select user
    db := mysql.Conn()
    user := entity.User{}
    result := db.Select("uid", "nickname").Take(&user, "uid = ?", touserid)
    if result.Error != nil {
        return nil, "user not exist"
    }

    //select followcount
    usercount := entity.UserCount{}
    result = db.Select("followcount", "followercount").Take(&usercount, "uid = ?", touserid)
    if result.Error != nil {
        usercount.Followcount = 0
        usercount.Followercount = 0
    }

    //selectisfollow
    var isfollow int8
    if userid != touserid {
        //userinfo owner or not,different user can see different userinfo
        ruserfollow := entity.RUserFollow{}
        result = db.Select("rufid").Take(&ruserfollow, "fromuid = ? and touid = ?", userid, touserid)
        if result.Error != nil {
            isfollow = NOTFOLLOW
        } else {
            isfollow = FOLLOWED
        }
    } else {
        isfollow = SAMEUSER
    }

    userinfo := UserInfo{
        ID: touserid,
        Name: user.Nickname,
        FollowCount: usercount.Followcount,
        FollowerCount: usercount.Followercount,
        IsFollow: isfollow,
    }

    returndata := map[string]interface{}{"user": userinfo}
    return returndata, ""
}
