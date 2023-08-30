package official_account_biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type SubscribeBizReq struct {
	AppId string

	ToUserName   string // 要推送的OpenId
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

// Subscribe 公众号关注事件
func (o *OfficialAccountUseCase) Subscribe(ctx context.Context, data SubscribeBizReq) (*SubscribeBizResp, error) {

	// 0.先判断该公众号是否在后台
	oaInfo, err := o.oaRepo.GetOfficialAccount(ctx, 0, data.AppId, "")
	if err != nil {
		o.log.Errorf("获取后台公众号配置信息失败：%v", err)
		return nil, err
	}

	// 1.新用户关注的时候，先去公众号拉取用户信息
	// 先获取accessToken
	info, err := o.oaRepo.RefreshAccessToken(ctx, data.AppId, oaInfo.AppSecret)
	if err != nil {
		o.log.Errorf("Subscribe o.oaRepo.GetAccessToken err:%v", err)
		return nil, err
	}

	// 2.获取公众号用户信息
	oaUserInfo, err := o.oaRepo.GetOAUserInfo(ctx, info.AccessToken, data.ToUserName)
	if err != nil {
		o.log.Errorf("Subscribe o.oaRepo.GetOAUserInfo err:%v", err)
		return nil, err
	}

	// 3.查看后台用户是否存在
	userInfo, err := o.oaRepo.GetUserInfo(ctx, data.AppId, "", oaUserInfo.UnionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 3.1 不存在就创
			_, err = o.oaRepo.CreateUserInfo(ctx, UserInfoDo{UnionID: userInfo.UnionID, OpenID: userInfo.OpenID})
			if err != nil {
				o.log.Errorf("Subscribe o.oaRepo.CreateUserInfo err:%v", err)
				return nil, err
			}
		} else {
			// 3.更新用户信息 没有就创建
			err = o.oaRepo.UpdateUserInfo(ctx, UserInfoDo{UnionID: userInfo.UnionID, OpenID: userInfo.OpenID})
			if err != nil {
				return nil, err
			}

		}
		return nil, err
	}

	// 4.普通关注
	var resp SubscribeBizResp
	resp.ToUserName = data.FromUserName
	resp.FromUserName = data.ToUserName
	resp.CreateTime = time.Now().Unix()
	resp.MsgType = "text"
	resp.Content = "欢迎关注"

	return &resp, nil
}

type SubscribeBizResp struct {
	ToUserName   string // 要推送的OpenId
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
