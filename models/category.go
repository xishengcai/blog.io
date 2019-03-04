package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	gorm.Model `json`
	Name string `gorm:"not null; size:12" json:"label"`
	CreteTime time.Time  `gorm:"not null" json:"-"`
	CategoryItem []CategoryItem `json:"categorys"`
}