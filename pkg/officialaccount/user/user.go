package user

import (
	"context"
	"encoding/json"
	"fmt"
	"wx-base/pkg/util"
)

const (
	userInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

// User 用户管理
type User struct {
}

// Info 用户基本信息
type Info struct {
	util.CommonError
	userInfo
}

// 用户基本信息
type userInfo struct {
	Subscribe      int32   `json:"subscribe"`
	OpenID         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int32   `json:"sex"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Province       string  `json:"province"`
	Language       string  `json:"language"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  int32   `json:"subscribe_time"`
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int32   `json:"groupid"`
	TagIDList      []int32 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int     `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
}

// GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(ctx context.Context, accessToken, openID string) (userInfo *Info, err error) {

	uri := fmt.Sprintf(userInfoURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(ctx, uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}
