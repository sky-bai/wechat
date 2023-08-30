package credential

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"wx-base/pkg/cache"
	"wx-base/pkg/util"
)

const (
	// accessTokenURL 获取access_token的接口
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// stableAccessTokenURL 获取稳定版access_token的接口
	stableAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/stable_token"

	// CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyOfficialAccountPrefix = "gowechat_officialaccount_"
	// CacheKeyMiniProgramPrefix 小程序cache key前缀
	CacheKeyMiniProgramPrefix = "gowechat_miniprogram_"
)

// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// StableAccessToken 获取稳定版接口调用凭据(与getAccessToken获取的调用凭证完全隔离，互不影响)
// 不强制更新access_token,可用于不同环境不同服务而不需要分布式锁以及公用缓存，避免access_token争抢
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
type StableAccessToken struct {
	appID          string
	appSecret      string
	cacheKeyPrefix string
	cache          cache.Cache
}

//

// GetAccessTokenDirectly 从微信获取access_token
func (d DefaultAccessToken) GetAccessTokenDirectly(ctx context.Context, appID, appSecret string, forceRefresh bool) (resAccessToken ResAccessToken, err error) {
	b, err := util.PostJSONContext(ctx, stableAccessTokenURL, map[string]interface{}{
		"grant_type":    "client_credential",
		"appid":         appID,
		"secret":        appSecret,
		"force_refresh": forceRefresh,
	})
	if err != nil {
		err = fmt.Errorf("util.PostJSONContext(%s) error : %s", stableAccessTokenURL, err)
		return
	}

	if err = json.Unmarshal(b, &resAccessToken); err != nil {
		err = fmt.Errorf("json.Unmarshal(%s) error : %s", b, err)
		return
	}

	if resAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get stable access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
