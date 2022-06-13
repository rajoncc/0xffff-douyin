package entity

import ()

type RUserFollow struct {
    Rufid uint64 `gorm:"primaryKey", json:"rufid"`
    Fromuid uint64 `json:"fromid"`
    Touid uint64 `json:"touid"`
    Isdel uint8 `json:"isdel"`
}

func (RUserFollow) TableName() string {
    return "r_user_follow"
}
