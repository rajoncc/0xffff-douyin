package entity

import ()

type RVideoFavorite struct {
    Rvfid uint64 `gorm:"primaryKey", json:"rvfid"`
    Uid uint64 `json:"uid"`
    Vid uint64 `json:"vid"`
    Isdel uint8 `json:"isdel"`
}

func (RVideoFavorite) TableName() string {
    return "r_video_favorite"
}
