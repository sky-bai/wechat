package auth

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"wx-base/pkg/officialaccount/wx_context"
	"wx-base/pkg/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataURL = "https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s"

	getPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
)

// Auth 登录/用户信息
type Auth struct {
	*wx_context.WxContext
}

// NewAuth new auth
func NewAuth(ctx *wx_context.WxContext) *Auth {
	return &Auth{ctx}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// RspCheckEncryptedData .
type RspCheckEncryptedData struct {
	util.CommonError

	Vaild      bool   `json:"vaild"`       // 是否是合法的数据
	CreateTime uint64 `json:"create_time"` // 加密数据生成的时间戳
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx context2.Context, appID, appSecret, jsCode string) (result ResCode2Session, err error) {
	var response []byte
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionURL, appID, appSecret, jsCode)); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// CheckEncryptedData .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedData(encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	return auth.CheckEncryptedDataContext(context2.Background(), encryptedMsgHash)
}

// CheckEncryptedDataContext .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedDataContext(ctx context2.Context, encryptedMsgHash string) (result RspCheckEncryptedData, err error) {
	var response []byte
	var (
		at string
	)
	if at, err = auth.GetAccessToken(); err != nil {
		return
	}

	// 由于GetPhoneNumberContext需要传入JSON，所以HTTPPostContext入参改为[]byte
	if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(checkEncryptedDataURL, at), []byte("encrypted_msg_hash="+encryptedMsgHash), nil); err != nil {
		return
	}
	if err = util.DecodeWithError(response, &result, "CheckEncryptedDataAuth"); err != nil {
		return
	}
	return
}

// GetPhoneNumberResponse 新版获取用户手机号响应结构体
type GetPhoneNumberResponse struct {
	util.CommonError

	PhoneInfo PhoneInfo `json:"phone_info"`
}

// PhoneInfo 获取用户手机号内容
type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
	WaterMark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"` // 数据水印
}

// GetPhoneNumberContext 小程序通过code获取用户手机号
func (auth *Auth) GetPhoneNumberContext(ctx context.Context, accessToken, code string) (*GetPhoneNumberResponse, error) {
	var response []byte
	var (
		err error
	)

	body := map[string]interface{}{
		"code": code,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	header := map[string]string{"Content-Type": "application/json;charset=utf-8"}
	if response, err = util.HTTPPostContext(ctx, fmt.Sprintf(getPhoneNumber, accessToken), bodyBytes, header); err != nil {
		return nil, err
	}

	var result GetPhoneNumberResponse
	if err = util.DecodeWithError(response, &result, "phonenumber.getPhoneNumber"); err != nil {
		return nil, err
	}
	return &result, nil
}
