package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wx-base/internal/biz"
	"wx-base/internal/data/model"
	"wx-base/pkg/officialaccount/oauth"
)

type OfficialAccountRepo struct {
	data *Data
	log  *log.Helper
}

func (o OfficialAccountRepo) CreateOfficialAccount(ctx context.Context, info biz.OfficialAccountInfo) (*biz.OfficialAccountInfo, error) {
	var data model.OfficialAccount
	data.AppID = info.AppID
	data.AppSecret = info.AppSecret
	data.Token = info.Token

	if err := o.data.db.Create(&model.OfficialAccount{}).Error; err != nil {
		return nil, err
	}

	return &biz.OfficialAccountInfo{
		Id:   data.ID,
		Name: data.Name,
	}, nil
}

// 这里dto 要转成 do

func (o OfficialAccountRepo) GetOpenidByCode(ctx context.Context, appId, code string) (*oauth.ResAccessToken, error) {
	//  1.根据appId获取appSecret
	var oa model.OfficialAccount
	if err := o.data.db.Debug().WithContext(ctx).Where("app_secret = ?", appId).First(&oa).Error; err != nil {
		return nil, err
	}

	// 2.根据appId和appSecret获取openid
	//credential.NewDefaultAccessToken(oa.AppID, oa.AppPwd, oa.Token, oa.AccessTokenExpire)
	data, err := o.data.oauth.GetUserAccessToken(ctx, appId, oa.AppSecret, code)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func NewOfficialAccountRepo(data *Data, logger log.Logger) biz.OfficialAccountRepo {
	return &OfficialAccountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetOfficialAccountByAppId 通过appId获取所有信息
func (o OfficialAccountRepo) GetOfficialAccountByAppId(ctx context.Context, appId string) (*biz.OfficialAccount, error) {
	o.data.db.Model(&biz.OfficialAccount{}).Where("app_id = ?", appId).First(&biz.OfficialAccount{})
	return nil, nil
}
