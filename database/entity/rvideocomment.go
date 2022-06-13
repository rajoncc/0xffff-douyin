package entity

import ()

type RVideoComment struct {
    Rvcid uint64 `gorm:"primaryKey", json:"rvcid"`
    Uid uint64 `json:"uid"`
    Vid uint64 `json:"vid"`
    Cid uint64 `json:"cid"`
    Isdel uint8 `json:"isdel"`
}

func (RVideoComment) TableName() string {
    return "r_video_comment"
}
