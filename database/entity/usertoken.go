package entity

import ()

type UserToken struct {
    Utid uint64 `gorm:"primaryKey", json:"utid"`
    Uid uint64 `json:"uid"`
    Token string `json:"token"`
    Expiredtime int64 `json:"expiredtime"`
}

func (UserToken) TableName() string {
    return "user_token"
}
