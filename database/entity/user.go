package entity

import ()

type User struct {
    Uid uint64 `gorm:"primaryKey", json:"uid"`
    Uname string `json:"uname"`
    Pword string `json:"pword"`
    Salt string `json:"salt"`
    Nickname string `json:"nickname"`
}

func (User) TableName() string {
    return "user"
}
