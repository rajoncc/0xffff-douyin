package entity

import ()

type Comment struct {
    Cid uint64 `gorm:"primaryKey", json:"cid"`
    Content string `json:"content"`
    Createtime int64 `json:"createtime"`
}

func (Comment) TableName() string {
    return "comment"
}
