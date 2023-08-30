package mini_program_service

import (
	"cdgitlib.spreadwin.cn/core/common/pb/pbWxBase"
	"context"
	"wx-base/internal/biz/mini_program_biz"
)

type MiniProgramService struct {
	pbWxBase.UnimplementedMiniProgramServer

	miniBiz *mini_program_biz.MiniProgramUseCase
}

func NewMiniProgramService(mini *mini_program_biz.MiniProgramUseCase) *MiniProgramService {
	return &MiniProgramService{miniBiz: mini}
}

// GetPhoneByCode 根据code获取手机号
func (m MiniProgramService) GetPhoneByCode(ctx context.Context, request *pbWxBase.GetMobileByCodeRequest) (*pbWxBase.GetMobileByCodeReply, error) {
	m.miniBiz.GetMobileByCode(ctx, mini_program_biz.GetMobileByCodeReq{AppId: request.AppId, Code: request.Code})
	return &pbWxBase.GetMobileByCodeReply{}, nil
}

func (m MiniProgramService) GetOpenidByCode(ctx context.Context, request *pbWxBase.GetOpenidByCodeRequest) (*pbWxBase.GetOpenidByCodeReply, error) {
	rv, err := m.miniBiz.GetOpenIdByCode(ctx, mini_program_biz.GetOpenIdByCodeReq{AppId: request.AppId, Customer: request.Customer, Code: request.Code})
	if err != nil {
		return nil, err
	}
	return &pbWxBase.GetOpenidByCodeReply{
		Data: &pbWxBase.GetOpenidByCodeReply_Data{
			Openid: rv.OpenID,
		},
	}, nil
}

func (m MiniProgramService) CreateMiniProgram(ctx context.Context, request *pbWxBase.CreateMiniProgramRequest) (*pbWxBase.CreateMiniProgramReply, error) {
	rv, err := m.miniBiz.CreateMiniProgram(ctx, mini_program_biz.MiniProgramInfo{
		AppId:     request.AppId,
		AppSecret: request.AppSecret,
		Customer:  request.Customer,
	})
	if err != nil {
		return nil, err
	}
	return &pbWxBase.CreateMiniProgramReply{
		Id: rv.Id,
	}, nil

}
