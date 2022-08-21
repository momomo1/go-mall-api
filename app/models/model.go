package models

import (
	"github.com/spf13/cast"
	"time"
)

//BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (base BaseModel) GetStringID() string {
	return cast.ToString(base.ID)
}
