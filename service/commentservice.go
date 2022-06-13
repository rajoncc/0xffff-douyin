package service

import (
    "time"

    "douyin/database/mysql"
    "douyin/database/entity"
)

type CommentInfo struct {
    ID uint64 `json:"id"`
    User map[string]interface{} `json:"user"`
    Content string `json:"content"`
    Createdate int64 `json:"create_date"`
}

type ReplyInfo struct {
    ID uint64 `json:"id"`
    Fromuid uint64 `json:"fromuid"`
    Fromname string `json:"fromnickname"`
    Touid uint64 `json:"touid"`
    Toname string `json:"tonickname"`
    Idtype uint8 `json:"idtype"`
    Content string `json:"content"`
    Createdate int64 `json:"create_date"`
}

func AddVideoComment(userid uint64, videoid uint64, content string) (map[string]interface{}, string) {
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

    //insert comment
    now := time.Now().Unix()
    comment := entity.Comment{Content: content, Createtime: now}
    result := tx.Create(&comment)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong comment"
    }

    //insert rvideocomment
    rvideocomment := entity.RVideoComment{Uid: userid, Vid: videoid, Cid: comment.Cid, Isdel: NOTDELETE}
    result = tx.Create(&rvideocomment)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong videos' comment"
    }

    //insert videocount
    videocount := entity.VideoCount{}
    result = tx.Select("commentcount").Take(&videocount, "vid = ?", videoid)
    if result.Error != nil {
        return nil, "video not exist"
    }
    
    videocount.Commentcount += 1
    result = tx.Model(&videocount).Where("vid = ?", videoid).Update("commentcount", videocount.Commentcount)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong videos' comment"
    }
    tx.Commit()

    //return comment[{id,user:userinfo,content,create_date}]
    userinfo, errmsg := GetUserInfo(userid, userid)
    if errmsg != "" {
        return nil, "user not exist"
    }

    commentinfo := CommentInfo{
        ID: comment.Cid,
        User: userinfo,
        Content: comment.Content,
        Createdate: comment.Createtime,
    }

    returndata := map[string]interface{}{"comment": commentinfo}
    return returndata, ""
}


func DelVideoComment(userid uint64, videoid uint64, commentid uint64) (map[string]interface{}, string) {
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

    //update rvideocomment
    rvideocomment := entity.RVideoComment{}
    result := tx.Model(&rvideocomment).Where("uid = ? and cid = ? and isdel = ?", userid, commentid, NOTDELETE).Update("isdel", DELETE)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "videos' comment not exist"
    }

    //update videocount
    videocount := entity.VideoCount{}
    result = tx.Select("commentcount").Take(&videocount, "vid = ?", videoid)
    if result.Error != nil {
        return nil, "video not exist"
    }
    if videocount.Commentcount >0 {
        videocount.Commentcount -= 1
    } else {
        videocount.Commentcount = 0
    }
    
    result = tx.Model(&videocount).Where("vid = ?", videoid).Update("commentcount", videocount.Commentcount)
    if result.Error != nil && result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong videos' comment"
    }

    tx.Commit()
    return nil, ""
}

func GetVideoCommentList(videoid uint64) (map[string]interface{}, string) {
    db := mysql.Conn()
    rvideocomments := []entity.RVideoComment{}
    result := db.Limit(COMMENTLISTLIMIT).Where("vid = ? and isdel = ?", videoid, NOTDELETE).Select("uid", "cid").Find(&rvideocomments)
    if result.Error != nil {
        return nil, "no comment"
    }
    
    commentlist := make([]CommentInfo, 0, COMMENTLISTLIMIT)
    for i := 0; i < len(rvideocomments); i++ {
        uid := rvideocomments[i].Uid
        cid := rvideocomments[i].Cid
        var commentinfo CommentInfo
        userinfo, errmsg := GetUserInfo(uid, uid)
        if errmsg != "" {
            return nil, errmsg
        }

        comment := entity.Comment{}
        result = db.Select("content", "createtime").Take(&comment, "cid = ?", cid)
        if result.Error != nil {
            return nil, "no comment"
        }

        commentinfo.ID = cid
        commentinfo.User = userinfo
        commentinfo.Content = comment.Content
        commentinfo.Createdate = comment.Createtime

        commentlist = append(commentlist, commentinfo)
    }
    
    returndata := map[string]interface{}{"comment_list": commentlist}
    return returndata, ""
}

func GetCommentReplyList(commentid uint64) (map[string]interface{}, string) {
    db := mysql.Conn()
    rcommentreplys := []entity.RCommentReply{}
    result := db.Limit(COMMENTLISTLIMIT).Where("cid = ? and isdel = ?", commentid, NOTDELETE).Select("cid", "fromid", "toid", "idtype").Find(&rcommentreplys)
    if result.Error != nil {
        return nil, "no reply"
    }
    
    replylist := make([]ReplyInfo, 0, COMMENTLISTLIMIT)
    for i := 0; i < len(rcommentreplys); i++ {
        fromid := rcommentreplys[i].Fromid
        toid := rcommentreplys[i].Toid
        idtype := rcommentreplys[i].Idtype
        var replyinfo ReplyInfo
        //fromid reply ruserreply user
        reply := entity.Reply{}
        result = db.Select("content", "createtime").Take(&reply, "rid = ?", fromid)
        if result.Error != nil {
            return nil, "no reply"
        }
        fromruserreply := entity.RUserReply{}
        result = db.Select("uid").Take(&fromruserreply, "rid = ?", fromid)
        if result.Error != nil {
            return nil, "no reply"
        }
        fromuser := entity.User{}
        result = db.Select("uid", "nickname").Take(&fromuser, "uid = ?", fromruserreply.Uid)
        if result.Error != nil {
            return nil, "no reply"
        }

        //if idtype toid ...
        touser := entity.User{Uid: 0, Nickname:""}
        if idtype == 1 {
            toruserreply := entity.RUserReply{}
            result = db.Select("uid").Take(&toruserreply, "rid = ?", toid)
            if result.Error != nil {
                return nil, "no reply"
            }
            touser := entity.User{}
            result = db.Select("uid", "nickname").Take(&touser, "uid = ?", toruserreply.Uid)
            if result.Error != nil {
                return nil, "no reply"
            }
        }

        replyinfo.ID = fromid
        replyinfo.Fromuid = fromuser.Uid
        replyinfo.Fromname = fromuser.Nickname
        replyinfo.Touid = touser.Uid
        replyinfo.Toname = touser.Nickname
        replyinfo.Idtype = idtype
        replyinfo.Content = reply.Content
        replyinfo.Createdate = reply.Createtime

        replylist = append(replylist, replyinfo)
    }
    
    returndata := map[string]interface{}{"reply_list": replylist}
    return returndata, ""
}

func AddCommentReply(userid uint64, commentid uint64, toid uint64, idtype uint8, content string) (map[string]interface{}, string) {
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

    //insert reply
    now := time.Now().Unix()
    reply := entity.Reply{Content: content, Createtime: now}
    result := tx.Create(&reply)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong reply"
    }

    //insert rcommentreply
    if idtype == 0 {
        toid = commentid
    }
    rcommentreply := entity.RCommentReply{Cid: commentid, Fromid: reply.Rid, Toid: toid,Idtype: idtype, Isdel: NOTDELETE}
    result = tx.Create(&rcommentreply)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong comments' reply"
    }

    ruserreply := entity.RUserReply{Uid: userid, Rid: reply.Rid}
    result = tx.Create(&ruserreply)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "wrong users' reply"
    }

    tx.Commit()

    return nil, ""
}

func DelCommentReply(fromid uint64) (map[string]interface{}, string) {
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

    //update rvideocomment
    rcommentreply := entity.RCommentReply{}
    result := tx.Model(&rcommentreply).Where("fromid = ? and isdel = ?", fromid, NOTDELETE).Update("isdel", DELETE)
    if result.Error != nil || result.RowsAffected < 1 {
        tx.Rollback()
        return nil, "comments' reply not exist"
    }

    tx.Commit()
    return nil, ""
}

