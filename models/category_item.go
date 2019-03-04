package models

import "time"

type CategoryItem struct{
	ID 	uint `gorm:"primary_key" json:"cid"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
	DeleteAt *time.Time `sql:"index" json:"-"`
	Name string `gorm:"not null; size:12" json:"cname"`
	CreateTime	time.Time `gorm:"not null" json:"create_time"`
	CategoryID uint `gorm:"index; not null" json:"-"`
	Article    []Article  `json:"-"`
	ItemSize   int        `gorm:"default:0" json:"size"`
	UserID     uint       `gorm:"index;not null" json:"-"`
}