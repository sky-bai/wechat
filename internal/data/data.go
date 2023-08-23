package data

import (
	"wx-base/internal/conf"
	"wx-base/pkg/miniprogram/auth"
	"wx-base/pkg/officialaccount/oauth"
	"wx-base/pkg/officialaccount/wx_context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewOauth, NewData, NewDB, NewMiniProgramRepo, NewOfficialAccountRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	oauth       *oauth.Oauth
	db          *gorm.DB
	miniProgram *auth.Auth
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

// NewOauth 实例化授权信息
func NewOauth(context *wx_context.WxContext) *oauth.Oauth {
	return oauth.NewOauth(context)
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// todo 这里还需要做具体的mysql配置
	return db
}
