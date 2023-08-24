package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "wx-base/api/wxbase/v1"
	"wx-base/internal/biz"
	"wx-base/pkg/util"
)

// OfficialAccountService is a OfficialAccount service.
type OfficialAccountService struct {
	v1.UnimplementedOfficialAccountServer

	uc *biz.OfficialAccountUseCase
}

// NewOfficialAccountService new a greeter service.
func NewOfficialAccountService(uc *biz.OfficialAccountUseCase) *OfficialAccountService {
	return &OfficialAccountService{uc: uc}
}

func (r *OfficialAccountService) GetOpenidByCode(ctx context.Context, in *v1.GetOpenidByCodeRequest) (*v1.GetOpenidByCodeReply, error) {
	data, err := r.uc.GetOpenidByCode(ctx, in.AppId, in.Code)
	if err != nil {
		return nil, err
	}
	return &v1.GetOpenidByCodeReply{
		Data: &v1.GetOpenidByCodeReply_Data{
			Openid: data.OpenID,
		},
	}, nil
}

func (r *OfficialAccountService) CreateOfficialAccount(ctx context.Context, in *v1.AddOfficialAccountRequest) (*v1.AddOfficialAccountReply, error) {

	rv, err := r.uc.CreateOfficialAccount(ctx, biz.OfficialAccountInfo{
		AppID:          in.AppId,
		AppSecret:      in.AppSecret,
		Token:          in.Token,
		EncodingAesKey: in.EncodingAesKey,
		AppIdAlias:     in.AppIdAlias,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AddOfficialAccountReply{
		Name: rv.Name,
		Id:   rv.Id,
	}, nil
}

// WxCallback 微信回调
func (r *OfficialAccountService) WxCallback(ctx context.Context, in *v1.WxCallbackRequest) (*v1.WxCallbackReply, error) {
	var info util.SignatureOptions
	info.Signature = in.Signature
	info.TimeStamp = in.Timestamp
	info.Nonce = in.Nonce
	info.EchoStr = in.Echostr

	if !util.Validate(info) {
		log.Errorf("签名验证失败")
		return nil, fmt.Errorf("签名验证失败")
	}

	return &v1.WxCallbackReply{
		Echostr: in.Echostr,
	}, nil

}
