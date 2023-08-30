package mini_program_biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type MiniProgramRepo interface {

	// CreateMiniProgram 小程序CRUD
	CreateMiniProgram(ctx context.Context, data MiniProgramInfo) (*MiniProgramInfo, error)
	GetMiniProgram(ctx context.Context, appId, customer string) (*MiniProgramInfo, error) // GetMiniProgram 获取小程序

	// GetOpenIdInfoByCode 根据code获取openId (或者是openId)
	GetOpenIdInfoByCode(ctx context.Context, data GetOpenIdInfoByCodeReq) (*GetOpenIdByCodeResp, error)
	GetMobileByCode(ctx context.Context, accessToken, code string) (*GetMobileByCodeResp, error)
	GetAccessToken(ctx context.Context, appId, appSecret string) (*AccessTokenInfo, error)
}
type MiniProgramUseCase struct {
	miniRepo MiniProgramRepo
	log      *log.Helper
}

func NewMiniProgramUseCase(repo MiniProgramRepo, logger log.Logger) *MiniProgramUseCase {
	return &MiniProgramUseCase{miniRepo: repo, log: log.NewHelper(logger)}
}

type MiniProgramInfo struct {
	Id             int64
	AppId          string
	Customer       string // 公众号英文名
	AppSecret      string
	JsToken        string
	Token          string
	ServerToken    string
	EncodingAesKey string
	AppIDAlias     string
}
type GetOpenIdByCodeReq struct {
	AppId     string
	Customer  string // 公众号英文名
	AppSecret string
	Code      string
}
type GetOpenIdInfoByCodeReq struct {
	AppId     string
	AppSecret string
	Code      string
}
type GetOpenIdByCodeResp struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}
type GetMobileByCodeReq struct {
	AppId    string
	Customer string // 公众号英文名
	Code     string
}
type GetMobileByCodeResp struct {
	Mobile string
}
type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (uc *MiniProgramUseCase) CreateMiniProgram(ctx context.Context, data MiniProgramInfo) (*MiniProgramInfo, error) {
	return uc.miniRepo.CreateMiniProgram(ctx, data)
}
func (uc *MiniProgramUseCase) GetMiniProgram(ctx context.Context, data MiniProgramInfo) (*MiniProgramInfo, error) {
	return nil, nil
}

func (uc *MiniProgramUseCase) GetOpenIdByCode(ctx context.Context, data GetOpenIdByCodeReq) (*GetOpenIdByCodeResp, error) {

	// 0.验证customer
	if data.AppId == "" && data.Customer == "" {
		return nil, errors.New("app_id or customer is required")
	}

	// 1.获取小程序配置
	miniInfo, err := uc.miniRepo.GetMiniProgram(ctx, data.AppId, data.Customer)
	if err != nil {
		return nil, err
	}

	// 2.根据code获取openId
	req := GetOpenIdInfoByCodeReq{}
	req.AppId = miniInfo.AppId
	req.AppSecret = miniInfo.AppSecret
	req.Code = data.Code
	openIdInfo, err := uc.miniRepo.GetOpenIdInfoByCode(ctx, req)
	if err != nil {
		return nil, err
	}

	return openIdInfo, nil
}
func (uc *MiniProgramUseCase) GetOpenIdInfoByCode(ctx context.Context, data GetOpenIdInfoByCodeReq) (*GetOpenIdByCodeResp, error) {
	return nil, nil
}
func (uc *MiniProgramUseCase) GetMobileByCode(ctx context.Context, data GetMobileByCodeReq) (*GetMobileByCodeResp, error) {

	// 0.根据appId获取小程序信息
	miniInfo, err := uc.miniRepo.GetMiniProgram(ctx, data.AppId, data.Customer)
	if err != nil {
		return nil, err
	}

	// 1.获取accessToken
	accInfo, err := uc.miniRepo.GetAccessToken(ctx, miniInfo.AppId, miniInfo.AppSecret)
	if err != nil {
		return nil, err
	}

	// 2.根据code获取手机号
	rv, err := uc.miniRepo.GetMobileByCode(ctx, accInfo.AccessToken, data.Code)
	if err != nil {
		return nil, err
	}

	return rv, nil

}
