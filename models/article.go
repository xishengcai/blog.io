package models

import "time"

type Article struct {
	ID  	uint `gorm:"primary_key" json:"aid"`
	CreateAt time.Time `jons:"-"`
	UpdateAt time.Time `json:"-"`
	DeleteAt *time.Time  `sql:"index" json:"-"`
	UserID   uint `gorm:"index;not null" json:"-"`
	CategoryItemID uint `gorm:"index;not null" json:"cid"`
	Title   string  `gorm:"type:varchar(100); not null" json:"title"`
	Content   string `gorm:"type:text; not null;" json:"content"`
	Cover    string  `gorm:"not null" json:"cover"`
	CreateTime time.Time `gorm:"not null" json:"create_time"`
	Views	int  `gorm:"default:0" json:"views"`
	Origin   int `gorm:"not null" json:"origin"`  // 是否是原创 1原创 0 转载
	State    int `gorm:"default:0" json:"-"`      // 0 正常发布	2并未发布(草稿箱)
}

func(ar Article) IsValid() bool {
	return ar.ID > 0
}