package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wx-base/pkg/officialaccount/oauth"
)

//业务逻辑的组装层，类似 DDD 的 domain 层

// OfficialAccount is a Greeter model.
type OfficialAccount struct {
	Hello string
}

// OfficialAccountRepo is a Greater repo.
type OfficialAccountRepo interface {
	GetOpenidByCode(ctx context.Context, appId, code string) (*oauth.ResAccessToken, error) // 只传appId和可选参数

	CreateOfficialAccount(ctx context.Context, info OfficialAccountInfo) (*OfficialAccountInfo, error)
}

// OfficialAccountUseCase is a Greeter usecase.
type OfficialAccountUseCase struct {
	repo OfficialAccountRepo
	log  *log.Helper
}

// NewOfficialAccountUseCase new a OfficialAccount usecase.
func NewOfficialAccountUseCase(repo OfficialAccountRepo, logger log.Logger) *OfficialAccountUseCase {
	return &OfficialAccountUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (o *OfficialAccountUseCase) GetOpenidByCode(ctx context.Context, appId, code string) (data *oauth.ResAccessToken, err error) {

	return o.repo.GetOpenidByCode(ctx, appId, code)
}

type OfficialAccountInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	Token          string `json:"token"`
	EncodingAesKey string `json:"encoding_aes_key"`
	AppIdAlias     string `json:"app_id_alias"`
}

// CreateOfficialAccount 创建公众号
func (o *OfficialAccountUseCase) CreateOfficialAccount(ctx context.Context, data OfficialAccountInfo) (*OfficialAccountInfo, error) {
	out, err := o.repo.CreateOfficialAccount(ctx, data)
	if err != nil {
		return nil, err
	}

	return out, nil
}
