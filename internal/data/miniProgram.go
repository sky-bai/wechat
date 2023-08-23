package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wx-base/internal/biz"
)

type miniProgramRepo struct {
	data *Data
	log  *log.Helper
}

func (m miniProgramRepo) GetMiniProgram(ctx context.Context, data biz.MiniProgramInfo) (biz.MiniProgramInfo, error) {
	var info biz.MiniProgramInfo
	info.AppId = data.AppId
	m.data.db.Create(&info)
	return info, nil
}

func (m miniProgramRepo) GetOpenIdByCode(ctx context.Context, data biz.GetOpenIdByCodeReq) (string, error) {
	sessionContext, err := m.data.miniProgram.Code2SessionContext(ctx, data.Code)
	if err != nil {
		return "", err
	}
	return sessionContext.OpenID, nil
}

func (m miniProgramRepo) CreateMiniProgram(ctx context.Context, info biz.MiniProgramInfo) error {
	//TODO implement me
	panic("implement me")
}

func NewMiniProgramRepo(data *Data, logger log.Logger) biz.MiniProgramRepo {
	return &miniProgramRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
