package midderware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"wx-base/internal/data"
)

// WxAuth 微信认证
func WxAuth(data *data.Data) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			//// 0.获取微信请求的参数
			//param := req.(*v1.WxCallbackRequest)
			//token := ""
			//
			//// 1.先判断是哪一个的请求 先获取到信息
			//if param.WxPlatform == "mini_program_service" {
			//	var oaInfo model.OfficialAccount
			//	if err := data.Db.Debug().WithContext(ctx).Where("app_id = ?", param.AppId).First(&oaInfo).Error; err != nil {
			//		return nil, err
			//	}
			//	token = oaInfo.Token
			//} else {
			//	var miniInfo model.MiniProgram
			//	if err := data.Db.Debug().WithContext(ctx).Where("app_id = ?", param.AppId).First(&miniInfo).Error; err != nil {
			//		return nil, err
			//	}
			//	token = miniInfo.Token
			//}
			//
			//// 2.先判断是否是微信服务器发来的请求
			//if param.Signature != util.Signature(token, param.Timestamp, param.Nonce) {
			//	return nil, errors.New("signature mismatch")
			//}

			reply, err = handler(ctx, req)

			return
		}
	}
}

// 1.获取到该小程序或者是公众号的配置信息
