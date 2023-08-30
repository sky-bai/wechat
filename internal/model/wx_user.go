package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameWxUser = "wx_user"

// WxUser mapped from table <wx_user>
type WxUser struct {
	ID           int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"` // 主键Id
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`          // 修改时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`          // 删除时间
	UnionID      string         `gorm:"column:union_id;type:varchar(255)" json:"union_id"`
	MiniPrograms []*MiniProgram
}

// TableName WxUser's table name
func (*WxUser) TableName() string {
	return TableNameWxUser
}
