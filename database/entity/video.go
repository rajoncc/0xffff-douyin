package entity

import ()

type Video struct {
    Vid uint64 `gorm:"primaryKey", json:"vid"`
    Title string `json:"title"`
    Playurl string `json:"playurl"`
    Coverurl string `json:"coverurl"`
    Createtime int64 `json:"createtime"`
}

func (Video) TableName() string {
    return "video"
}
