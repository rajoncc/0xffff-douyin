package entity

import ()

type RUserReply struct {
    Rurid uint64 `gorm:"primaryKey", json:"rurid"`
    Uid uint64 `json:"uid"`
    Rid uint64 `json:"rid"`
}

func (RUserReply) TableName() string {
    return "r_user_reply"
}

