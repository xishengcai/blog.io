package models

import "time"

type FriendlyLink struct {
	ID uint `gorm:"primary_key" json:"link_id"` // 自增
	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"-"`
	UserID    uint      `gorm:"index;not null" json:"-"`
	Avator    string    `gorm:"not null" json:"link_icon"`
	LinkName  string    `gorm:"not null;size:12" json:"link_name"`
	LinkUrl   string    `gorm:"not null" json:"link_url"`
}