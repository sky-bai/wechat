package offiaiacl_account_data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"wx-base/internal/biz/official_account_biz"
	"wx-base/internal/data"
	"wx-base/internal/model"
	"wx-base/pkg/officialaccount/oauth"
)

var ProviderSet = wire.NewSet(NewOfficialAccountRepo)

var _ official_account_biz.OfficialAccountRepo = (*officialAccountRepo)(nil)

type officialAccountRepo struct {
	data *data.Data
	log  *log.Helper
}

// RefreshAccessToken 重新刷新并获取公众号的access_token
func (o *officialAccountRepo) RefreshAccessToken(ctx context.Context, appId, appSecret string) (*official_account_biz.AccessTokenInfo, error) {
	acInfo, err := o.data.OfficialAccount.AccessToken.GetAccessTokenDirectly(ctx, appId, appSecret, true)
	if err != nil {
		o.log.Errorf("获取公众号access_token失败：%v", err)
		return nil, err
	}

	return &official_account_biz.AccessTokenInfo{
		AccessToken: acInfo.AccessToken,
		ExpiresIn:   acInfo.ExpiresIn,
	}, nil
}

func NewOfficialAccountRepo(data *data.Data, logger log.Logger) official_account_biz.OfficialAccountRepo {
	return &officialAccountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o *officialAccountRepo) Subscribe(ctx context.Context, appId, openId string) error {
	//TODO implement me
	panic("implement me")
}

// GetOAUserInfo 从公众号获取用户信息
func (o *officialAccountRepo) GetOAUserInfo(ctx context.Context, accessToken, openId string) (*official_account_biz.UserInfoDo, error) {
	oaUserInfo, err := o.data.OfficialAccount.User.GetUserInfo(ctx, accessToken, openId)
	if err != nil {
		return nil, err
	}
	return &official_account_biz.UserInfoDo{
		UnionID:   oaUserInfo.UnionID,
		Subscribe: oaUserInfo.Subscribe,
		OpenID:    oaUserInfo.OpenID,
	}, nil
}

// CreateUserInfo 后台用户CRUD
func (o *officialAccountRepo) CreateUserInfo(ctx context.Context, userInfo official_account_biz.UserInfoDo) (int64, error) {
	var wxUserData model.WxUser
	wxUserData.UnionID = userInfo.UnionID
	err := o.data.Db.Debug().WithContext(ctx).Create(&wxUserData).Error
	if err != nil {
		return 0, err
	}
	// 添加中间表数据
	o.data.Db.Debug().WithContext(ctx).Model(&wxUserData).Association("MiniPrograms").Append(&model.MiniProgram{})
	return wxUserData.ID, nil
}
func (o *officialAccountRepo) GetUserInfo(ctx context.Context, appId, id, unionId string) (*official_account_biz.UserInfoDo, error) {

	var data model.WxUser
	err := o.data.Db.Debug().WithContext(ctx).Where("union_id = ? OR id = ? ", unionId, id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &official_account_biz.UserInfoDo{
		Id: data.ID,
	}, nil
}
func (o *officialAccountRepo) UpdateUserInfo(ctx context.Context, userInfo official_account_biz.UserInfoDo) error {
	var data model.WxUser
	data.UnionID = userInfo.UnionID
	err := o.data.Db.Debug().WithContext(ctx).Save(&data).Error
	if err != nil {
		return err
	}

	return nil
}
func (o *officialAccountRepo) DeleteUserInfo(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

// CreateOfficialAccount 公众号CRUD
func (o *officialAccountRepo) CreateOfficialAccount(ctx context.Context, info official_account_biz.OfficialAccountInfo) (int64, error) {
	var oaData model.OfficialAccount
	oaData.AppID = info.AppID
	oaData.AppSecret = info.AppSecret
	oaData.Token = info.Token
	oaData.EncodingAesKey = info.EncodingAesKey
	oaData.AppIDAlias = info.AppIdAlias
	oaData.Name = info.Name
	oaData.Customer = info.Customer
	oaData.Url = info.Url

	if err := o.data.Db.Debug().WithContext(ctx).Create(&oaData).Error; err != nil {
		return 0, err
	}

	return oaData.ID, nil
}
func (o *officialAccountRepo) GetOfficialAccount(ctx context.Context, id int64, appId, customer string) (*official_account_biz.OfficialAccountInfo, error) {
	var oaInfo model.OfficialAccount
	db := o.data.Db.Debug().WithContext(ctx)
	if id != 0 {
		db = db.Where("id = ?", id)
	}
	if appId != "" {
		db = db.Where("app_id = ?", appId)
	}
	if customer != "" {
		db = db.Where("customer = ?", customer)
	}

	err := db.First(&oaInfo).Error
	if err != nil {
		return nil, err
	}

	return &official_account_biz.OfficialAccountInfo{
		Id:             oaInfo.ID,
		Name:           oaInfo.Name,
		AppID:          oaInfo.AppID,
		AppSecret:      oaInfo.AppSecret,
		Token:          oaInfo.Token,
		EncodingAesKey: oaInfo.EncodingAesKey,
		AppIdAlias:     oaInfo.AppIDAlias,
	}, nil
}
func (o *officialAccountRepo) UpdateOfficialAccount(ctx context.Context) (*official_account_biz.OfficialAccountInfo, error) {
	//TODO implement me
	panic("implement me")
}
func (o *officialAccountRepo) DeleteOfficialAccount(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

// SendTemplateMessage 发送模板消息
func (o *officialAccountRepo) SendTemplateMessage(ctx context.Context, appId, openId, templateId string, data map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

// GetAccessToken 获取access_token
func (o *officialAccountRepo) GetAccessToken(ctx context.Context, appId, appSecret string) (*official_account_biz.AccessTokenInfo, error) {

	// 1.根据appId和appSecret获取token 直接从数据库中获取
	var oaInfo model.OfficialAccount
	err := o.data.Db.Debug().WithContext(ctx).Select("access_token, expire_in").Where("app_id = ? and app_secret = ?", appId, appSecret).First(&oaInfo).Error
	if err != nil {
		return nil, err
	}

	return &official_account_biz.AccessTokenInfo{
		AccessToken: oaInfo.AccessToken,
	}, nil

}

func (o *officialAccountRepo) GetOpenidByCode(ctx context.Context, appId, code string) (*oauth.ResAccessToken, error) {
	//  1.根据appId获取appSecret
	var oa model.OfficialAccount
	if err := o.data.Db.Debug().WithContext(ctx).Where("app_secret = ?", appId).First(&oa).Error; err != nil {
		return nil, err
	}

	// 2.根据appId和appSecret获取openid
	//credential.NewDefaultAccessToken(oa.AppID, oa.AppPwd, oa.Token, oa.AccessTokenExpire)
	//data, err := o.data.Oauth.GetUserAccessToken(ctx, appId, oa.AppSecret, code)
	//if err != nil {
	//	return nil, err
	//}
	//return &data, nil
	return nil, nil
}

// GetOfficialAccountByAppId 通过appId获取所有信息
func (o *officialAccountRepo) GetOfficialAccountByAppId(ctx context.Context, appId string) (*official_account_biz.OfficialAccount, error) {
	o.data.Db.Model(&official_account_biz.OfficialAccount{}).Where("app_id = ?", appId).First(&official_account_biz.OfficialAccount{})
	return nil, nil
}
