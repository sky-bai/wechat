package officialaccount

import (
	"wx-base/pkg/credential"
	"wx-base/pkg/officialaccount/message"
	"wx-base/pkg/officialaccount/oauth"
	"wx-base/pkg/officialaccount/user"
	"wx-base/pkg/officialaccount/wx_context"
)

// OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx         *wx_context.WxContext
	Oauth       *oauth.Oauth
	TemplateMsg *message.Template
	User        *user.User
	AccessToken credential.DefaultAccessToken
}
