package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type MiniProgramRepo interface {
	// CreateMiniProgram 创建小程序
	CreateMiniProgram(ctx context.Context, data MiniProgramInfo) error
	// GetMiniProgram 获取小程序
	GetMiniProgram(ctx context.Context, data MiniProgramInfo) (MiniProgramInfo, error)

	// GetOpenIdByCode 根据code获取openId
	GetOpenIdByCode(ctx context.Context, data GetOpenIdByCodeReq) (string, error)
}

type MiniProgramUseCase struct {
	miniRepo MiniProgramRepo
	log      *log.Helper
}

func NewMiniProgramUseCase(repo MiniProgramRepo, logger log.Logger) *MiniProgramUseCase {
	return &MiniProgramUseCase{miniRepo: repo, log: log.NewHelper(logger)}
}

type MiniProgramInfo struct {
	AppId string `json:"app_id"`
}

type GetOpenIdByCodeReq struct {
	Code    string
	AppId   string
	AppName string
}

func (uc *MiniProgramUseCase) CreateMiniProgram(ctx context.Context, data MiniProgramInfo) error {
	return uc.miniRepo.CreateMiniProgram(ctx, data)
}

func (uc *MiniProgramUseCase) GetOpenIdByCode(ctx context.Context, data GetOpenIdByCodeReq) (string, error) {

	// 0.验证customer
	if data.AppId == "" && data.AppName == "" {
		return "", fmt.Errorf("appId and appName can not be empty")
	}

	// 1.获取小程序配置

	// 2.根据code获取openId
	uc.miniRepo.GetOpenIdByCode(ctx, data)

	// 3.新建小程序用户账号

	return "", nil
}
