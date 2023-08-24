package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	v1 "wx-base/api/wxbase/v1"
	"wx-base/internal/conf"
	"wx-base/internal/service"
	"wx-base/pkg/midderware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, oa *service.OfficialAccountService, mini *service.MiniProgramService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			midderware.WxAuth(),
			//	selector.Server(
			//	// 这里面填具体的中间件处理逻辑
			//	).Match(
			//		// 这里面写需要匹配的接口
			//		NewWhiteListMatcher()).
			//		Build(),
			//),
			selector.Server(recovery.Recovery(), midderware.WxAuth()).
				Path("/wx/new/new").
				//Prefix("/wx/new/new").
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	// openApI
	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)

	// 微信发过来的消息
	srv.HandleFunc("/wx/new", service.WxBase())

	// 单独起的http服务好像没加入中间件昵

	//v1.RegisterWxBaseHTTPServer()

	v1.RegisterOfficialAccountHTTPServer(srv, oa)
	v1.RegisterMiniProgramHTTPServer(srv, mini)
	return srv
}

//
//func NewWhiteListMatcher() selector.MatchFunc {
//
//	whiteList := make(map[string]struct{})
//	whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
//	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
//	return func(ctx context.Context, operation string) bool {
//		if _, ok := whiteList[operation]; ok {
//			return false
//		}
//		return true
//	}
//}
