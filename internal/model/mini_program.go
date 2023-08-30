package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMiniProgram = "mini_program"

// MiniProgram mapped from table <mini_program>
type MiniProgram struct {
	ID              int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"` // 主键Id
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`          // 修改时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`          // 删除时间
	Customer        string         `gorm:"column:customer;type:varchar(255)" json:"customer"`
	Name            string         `gorm:"column:name;type:varchar(255)" json:"name"`
	AppID           string         `gorm:"column:app_Id;type:varchar(255)" json:"app_Id"`
	AppSecret       string         `gorm:"column:app_secret;type:varchar(255)" json:"app_secret"`
	JsToken         string         `gorm:"column:js_token;type:varchar(255)" json:"js_token"`
	Token           string         `gorm:"column:token;type:varchar(255)" json:"token"`
	TokenUpdateTime *time.Time     `gorm:"column:token_update_time;type:datetime(3)" json:"token_update_time"`
	URL             string         `gorm:"column:url;type:varchar(255)" json:"url"`
	ServerToken     string         `gorm:"column:server_token;type:varchar(255)" json:"server_token"`
	EncodingAesKey  string         `gorm:"column:encoding_aes_key;type:varchar(255)" json:"encoding_aes_key"`
	AppIDAlias      string         `gorm:"column:app_Id_alias;type:varchar(255);comment:appId的别名" json:"app_Id_alias"` // appId的别名
	WxUsers         []*WxUser
}

// TableName MiniProgram's table name
func (*MiniProgram) TableName() string {
	return TableNameMiniProgram
}
