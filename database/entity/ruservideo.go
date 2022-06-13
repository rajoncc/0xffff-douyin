package entity

import ()

type RUserVideo struct {
    Ruvid uint64 `gorm:"primaryKey", json:"ruvid"`
    Uid uint64 `json:"uid"`
    Vid uint64 `json:"vid"`
    Isdel uint8 `json:"isdel"`
}

func (RUserVideo) TableName() string {
    return "r_user_video"
}
