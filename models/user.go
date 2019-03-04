package models

import (
	"database/sql"
	"time"
)

type User struct {
	BaseModel

	UserName     string `gorm:"not null;size:12;unique"`
	UserPassWord string `gorm:"not null;size:40"`
	LoginIP      sql.NullString
	LoginTime    time.Time
	IsActive     int      `gorm:"not null; default: 0"`
	UserInfo     UserInfo `gorm:"ForeignKey:UserInfoID; AssociationForeignKey:ID"`
	UserInfoID   uint     `gorm:"index; not null"`
	Article      []Article
}
