package data

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm/schema"
	"wx-base/internal/conf"
	"wx-base/internal/model"
	"wx-base/pkg/miniprogram"
	"wx-base/pkg/miniprogram/auth"
	"wx-base/pkg/officialaccount"
	"wx-base/pkg/officialaccount/oauth"
	"wx-base/pkg/officialaccount/wx_context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewRedisCmd)

// Data .
type Data struct {
	// TODO wrapped database client
	Db  *gorm.DB
	rdb *redis.Client

	//Oauth       *oauth.Oauth
	MiniProgram *auth.Auth
	// 这里提供
	Mini            *miniprogram.ThirdMiniProgram
	OfficialAccount officialaccount.OfficialAccount
}

// NewData .
func NewData(gormClient *gorm.DB, redisCli *redis.Client, logger log.Logger) (*Data, func(), error) {

	d := &Data{
		Db:  gormClient,
		rdb: redisCli,
	}

	return d, func() {
		log.NewHelper(logger).Info("closing the data resources")

		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// NewOauth 实例化授权信息
func NewOauth(context *wx_context.WxContext) *oauth.Oauth {
	return oauth.NewOauth(context)
}

func NewGormDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "service/data/gorm"))
	// 1.连接gorm
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		SkipDefaultTransaction: false, // 是否以事务模式开启数据库操作
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true, // skip the snake_casing of names
		},
		DisableForeignKeyConstraintWhenMigrating: true, //  禁用物理外建使用逻辑外建
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2.创建表
	err = db.AutoMigrate(&model.OfficialAccount{})
	if err != nil {
		log.Fatalf("failed to AutoMigrate : %v", err)
	}

	return db
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) *redis.Client {
	//log := log.NewHelper(log.With(logger, "module", "user-service/data/redis"))
	// redis
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	//timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	//defer cancelFunc()
	//err := client.Ping(timeout).Err()
	//if err != nil {
	//	log.Fatalf("redis connect error: %v", err)
	//}
	return client
}
