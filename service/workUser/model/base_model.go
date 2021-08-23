package model

import (
	"gorm.io/gorm"
	"time"
)

// 参考 gorm.Model
type BaseModel struct {
	ID        uint64         `gorm:"primarykey;comment:主键"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}
