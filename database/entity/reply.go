package entity

import ()

type Reply struct {
    Rid uint64 `gorm:"primaryKey", json:"rid"`
    Content string `json:"content"`
    Createtime int64 `json:"createtime"`
}

func (Reply) TableName() string {
    return "reply"
}
