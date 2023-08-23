package wx_context

import (
	"wx-base/pkg/credential"
	"wx-base/pkg/officialaccount/config"
)

// WxContext struct
type WxContext struct {
	*config.Config
	credential.AccessTokenHandle
}
