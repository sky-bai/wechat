package model

import (
	"gorm.io/gorm"
	"time"
)

// 自定义连接表

type MiniUser struct {
	ID            int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"` // 主键Id
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`          // 修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`          // 删除时间
	WxUserID      int64          `gorm:"column:wx_user_id;primaryKey"`
	MiniProgramID int64          `gorm:"column:mini_program_id;primaryKey"`
	Mobile        string         `gorm:"column:mobile;type:varchar(11)" json:"mobile"`
	OpenId        string         `gorm:"column:open_id;type:varchar(255)" json:"open_id"`

	JsToken string `gorm:"column:js_token;type:varchar(255)" json:"js_token"`
}

// TableName MiniProgram's table name
func (*MiniUser) TableName() string {
	return "mini_user"
}
