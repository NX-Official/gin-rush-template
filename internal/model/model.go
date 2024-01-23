package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Model) CreateTime() int64 {
	return m.CreatedAt.UnixMilli()
}

func (m *Model) UpdateTime() int64 {
	return m.UpdatedAt.UnixMilli()
}

type Dto struct {
	ID         uint  `json:"id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}
