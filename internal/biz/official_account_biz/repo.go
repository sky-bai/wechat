package official_account_biz

import (
	"context"
	"wx-base/pkg/officialaccount/oauth"
)

// OfficialAccountRepo is a Greater repo.
type OfficialAccountRepo interface {
	// Subscribe 公众号关注事件
	Subscribe(ctx context.Context, appId, openId string) error

	GetOpenidByCode(ctx context.Context, appId, code string) (*oauth.ResAccessToken, error) // 只传appId和可选参数
	GetAccessToken(ctx context.Context, appId, appSecret string) (*AccessTokenInfo, error)
	RefreshAccessToken(ctx context.Context, appId, appSecret string) (*AccessTokenInfo, error)
	SendTemplateMessage(ctx context.Context, appId, openId, templateId string, data map[string]interface{}) error

	// GetOAUserInfo 获取公众号的用户信息
	GetOAUserInfo(ctx context.Context, accessToken, openId string) (*UserInfoDo, error)

	// CreateUserInfo 公众号下用户信息CRUD
	CreateUserInfo(ctx context.Context, data UserInfoDo) (int64, error)
	GetUserInfo(ctx context.Context, appId, id, unionId string) (*UserInfoDo, error)
	UpdateUserInfo(ctx context.Context, data UserInfoDo) error
	DeleteUserInfo(ctx context.Context) error

	// CreateOfficialAccount 公众号CRUD
	CreateOfficialAccount(ctx context.Context, info OfficialAccountInfo) (int64, error)
	GetOfficialAccount(ctx context.Context, id int64, appId, customer string) (*OfficialAccountInfo, error)
	UpdateOfficialAccount(ctx context.Context) (*OfficialAccountInfo, error)
	DeleteOfficialAccount(ctx context.Context) error
}

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type UserInfoDo struct {
	Id             int64   `json:"id"`
	Subscribe      int32   `json:"subscribe"` // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenID         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int32   `json:"sex"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Province       string  `json:"province"`
	Language       string  `json:"language"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  int32   `json:"subscribe_time"` // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int32   `json:"groupid"`
	TagIDList      []int32 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int     `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
}
