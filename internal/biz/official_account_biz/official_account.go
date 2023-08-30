package official_account_biz

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

// OfficialAccountUseCase is a Greeter usecase.
type OfficialAccountUseCase struct {
	oaRepo OfficialAccountRepo
	log    *log.Helper
}

func (o *OfficialAccountUseCase) RefreshAccessToken(ctx context.Context, appId string) (*AccessTokenInfo, error) {
	oaInfo, err := o.oaRepo.GetOfficialAccount(ctx, 0, appId, "")
	if err != nil {
		return nil, err
	}

	return o.oaRepo.RefreshAccessToken(ctx, appId, oaInfo.AppSecret)
}

func (o *OfficialAccountUseCase) GetAccessToken(ctx context.Context, appId string) (data *AccessTokenInfo, err error) {

	// 1.根据appID获取appSecret
	oaInfo, err := o.oaRepo.GetOfficialAccount(ctx, 0, appId, "")
	if err != nil {
		return nil, err
	}

	// 2.根据appID和appSecret获取token
	rv, err := o.oaRepo.GetAccessToken(ctx, appId, oaInfo.AppSecret)
	if err != nil {
		return nil, err
	}

	return rv, nil
}

// NewOfficialAccountUseCase new a OfficialAccount usecase.
func NewOfficialAccountUseCase(repo OfficialAccountRepo, logger log.Logger) *OfficialAccountUseCase {
	return &OfficialAccountUseCase{oaRepo: repo, log: log.NewHelper(logger)}
}

func (o *OfficialAccountUseCase) GetOpenidByCode(ctx context.Context, appId, code string) (data *oauth.ResAccessToken, err error) {

	return o.oaRepo.GetOpenidByCode(ctx, appId, code)
}

type OfficialAccountInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`       // 公众号中文名
	AppID          string `json:"app_id"`     // 公众号appId 行车云助手
	Customer       string `json:"customer"`   // 公众号别名 XingCheYunZhuShou
	AppSecret      string `json:"app_secret"` // 公众号appSecret
	Token          string `json:"token"`      // 微信令牌
	EncodingAesKey string `json:"encoding_aes_key"`
	AppIdAlias     string `json:"app_id_alias"`
	Url            string `json:"url"` // 微信回调地址
}

// CreateOfficialAccount 公众号CRUD
func (o *OfficialAccountUseCase) CreateOfficialAccount(ctx context.Context, data OfficialAccountInfo) (int64, error) {
	out, err := o.oaRepo.CreateOfficialAccount(ctx, data)
	if err != nil {
		return 0, err
	}

	return out, nil
}
func (o *OfficialAccountUseCase) GetOfficialAccount(ctx context.Context, id int64, appId, customer string) (*OfficialAccountInfo, error) {
	o.oaRepo.GetOfficialAccount(ctx, id, appId, customer)
	return nil, nil
}

type SendTemplateMessage struct {
	TempType            string // 模板别名（与templateId二选一）
	TemplateId          string // 模板ID（与tempType二选一）
	AppId               string // 要发往的公众号appId(与customer二选一)
	Customer            string // 要发往的公众号别名(与appId二选一)
	OpenId              string // 接收者openid（选填）
	Mobile              string // 接收者手机号（选填）
	UnionId             string // 接收者unionId（选填）
	Url                 string // 模板跳转链接（选填）
	MiniProgramAppId    string // 跳转小程序的appid
	MiniProgramPagePath string // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data                string // 模板数据 json格式
}

// SendTemplateMessage 发送模板消息
func (o *OfficialAccountUseCase) SendTemplateMessage(ctx context.Context, data SendTemplateMessage) error {
	// 0.如果用户已经取关了，就不发模板消息了
	//

	// 1.最后发送模版消息
	//o.repo.SendTemplateMessage()

	return nil
}
