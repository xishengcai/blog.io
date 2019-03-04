package models

import (
	"blogserver/common"
	"time"
)

type BaseModel struct {
	ID uint `gorm:"primary_key"`  // 自增
	CreateAt  time.Time  `json:"-"`
	UpdateAt  time.Time   `json:"-"`
}

// Api 相应结构体
type Response struct {
	common.Err
	Data interface{} `json:"data"`
}

type SortValue struct {
	Key string
	Value string
}

func (sv SortValue) IsValidValue() bool {
	return sv.Value == "desc" || sv.Value == "asc"
}

func (sv SortValue) IsValidKey() bool {
	return sv.Key == "create_time"  || sv.Key == "views"
}