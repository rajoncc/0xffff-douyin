package entity

import ()

type VideoCount struct {
    Vcid uint64 `gorm:"primaryKey", json:"vcid"`
    Vid uint64 `json:"vid"`
    Favoritecount uint32 `json:"favoritecount"`
    Commentcount uint32 `json:"commentcount"`
}

func (VideoCount) TableName() string {
    return "video_count"
}
