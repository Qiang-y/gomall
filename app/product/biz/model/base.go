package model

import "time"

type Base struct {
	ID       int `gorm:"primaryKey"`
	CreateAt time.Time
	UpdataAt time.Time
}
