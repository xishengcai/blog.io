package models

type UserInfo struct {
	BaseModel
	UserAvator string `gorm:"not null; default:''"`
	UserDesc   string `gorm:"not null;"`
	UserEmail  string `gorm:"not null"`
	UserAddr   string `gorm:"not null"`
	NickName   string `gorm:"not null;size:12;"`
}