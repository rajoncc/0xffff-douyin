package entity

import ()

type RCommentReply struct {
    Rcrid uint64 `gorm:"primaryKey", json:"rcrid"`
    Cid uint64 `json:"cid"`
    Fromid uint64 `json:"fromid"`
    Toid uint64 `json:"toid"`
    Idtype uint8 `json:"idtype"`
    Isdel uint8 `json:"isdel"`
}

func (RCommentReply) TableName() string {
    return "r_comment_reply"
}
