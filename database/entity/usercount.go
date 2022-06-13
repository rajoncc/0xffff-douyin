package entity

import ()

type UserCount struct {
    Ucid uint64 `gorm:"primaryKey", json:"ucid"`
    Uid uint64 `json:"uid"`
    Followcount uint32 `json:"followcount"`
    Followercount uint32 `json:"followercount"`
}

func (UserCount) TableName() string {
    return "user_count"
}
