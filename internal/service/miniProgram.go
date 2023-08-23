package service

import (
	"context"
	v1 "wx-base/api/wxbase/v1"
	"wx-base/internal/biz"
)

// MiniProgramService is a MiniProgram service.
type MiniProgramService struct {
	v1.UnimplementedMiniProgramServer

	uc *biz.MiniProgramUseCase
}

// NewMiniProgramService new a greeter service.
func NewMiniProgramService(uc *biz.MiniProgramUseCase) *MiniProgramService {
	return &MiniProgramService{uc: uc}
}

func (m *MiniProgramService) GetOpenidByCode(ctx context.Context, in *v1.GetOpenidByCodeRequest) (*v1.GetOpenidByCodeReply, error) {
	code, err := m.uc.GetOpenIdByCode(ctx, biz.GetOpenIdByCodeReq{Code: in.Code, AppId: in.AppId})
	if err != nil {
		return nil, err
	}

	return &v1.GetOpenidByCodeReply{
		Data: &v1.GetOpenidByCodeReply_Data{
			Openid: code,
		},
	}, nil

}
