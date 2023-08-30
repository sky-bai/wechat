package mini_program_data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"wx-base/internal/biz/mini_program_biz"
	"wx-base/internal/data"
	"wx-base/internal/model"
)

var ProviderSet = wire.NewSet(NewMiniProgramRepo)

var _ mini_program_biz.MiniProgramRepo = (*miniProgramRepo)(nil)

type miniProgramRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMiniProgramRepo(data *data.Data, logger log.Logger) mini_program_biz.MiniProgramRepo {
	return &miniProgramRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m miniProgramRepo) GetMiniProgram(ctx context.Context, appId, customer string) (*mini_program_biz.MiniProgramInfo, error) {
	var info model.MiniProgram
	err := m.data.Db.Debug().WithContext(ctx).Where("app_id = ? OR customer = ?", appId, customer).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &mini_program_biz.MiniProgramInfo{
		Id:        info.ID,
		AppId:     info.AppID,
		AppSecret: info.AppSecret,
	}, nil
}

func (m miniProgramRepo) GetAccessToken(ctx context.Context, appId, appSecret string) (*mini_program_biz.AccessTokenInfo, error) {
	//TODO implement me
	panic("implement me")
}

// CreateMiniProgram 小程序CRUD
func (m miniProgramRepo) CreateMiniProgram(ctx context.Context, data mini_program_biz.MiniProgramInfo) (*mini_program_biz.MiniProgramInfo, error) {
	var miniData model.MiniProgram
	miniData.AppID = data.AppId
	miniData.AppSecret = data.AppSecret
	miniData.Customer = data.Customer
	miniData.Token = data.Token

	err := m.data.Db.Create(&miniData).Error
	if err != nil {
		return nil, err
	}

	return &mini_program_biz.MiniProgramInfo{
		Id: miniData.ID,
	}, nil
}

// GetOpenIdInfoByCode 通过code获取openid
func (m miniProgramRepo) GetOpenIdInfoByCode(ctx context.Context, data mini_program_biz.GetOpenIdInfoByCodeReq) (*mini_program_biz.GetOpenIdByCodeResp, error) {

	sessionContext, err := m.data.MiniProgram.Code2SessionContext(ctx, data.AppId, data.AppSecret, data.Code)
	if err != nil {
		return nil, err
	}
	return &mini_program_biz.GetOpenIdByCodeResp{
		OpenID:     sessionContext.OpenID,
		SessionKey: sessionContext.SessionKey,
		UnionID:    sessionContext.UnionID,
	}, nil
}

// GetMobileByCode 通过code获取手机号
func (m miniProgramRepo) GetMobileByCode(ctx context.Context, accessToken, code string) (*mini_program_biz.GetMobileByCodeResp, error) {
	rv, err := m.data.MiniProgram.GetPhoneNumberContext(ctx, accessToken, code)
	if err != nil {
		return nil, err
	}
	return &mini_program_biz.GetMobileByCodeResp{
		Mobile: rv.PhoneInfo.PhoneNumber,
	}, nil
}
