package server

import (
	"cdgitlib.spreadwin.cn/core/common/pb/pbWxBase"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	nethttp "net/http"
	"strings"
	"wx-base/internal/conf"
	"wx-base/internal/data"
	"wx-base/internal/service/mini_program_service"
	"wx-base/internal/service/official_account_service"
	"wx-base/pkg/midderware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, data *data.Data, oa *official_account_service.OfficialAccountService, mini *mini_program_service.MiniProgramService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ResponseEncoder(CustomerResponseEncoder),
		http.Middleware(
			validate.Validator(),
			midderware.WxAuth(data),
			//	selector.Server(
			//	// 这里面填具体的中间件处理逻辑
			//	).Match(
			//		// 这里面写需要匹配的接口
			//		NewWhiteListMatcher()).
			//		Build(),
			//),
			//selector.Server(recovery.Recovery(), midderware.WxAuth()).
			//	Path("/wx/new/new").
			//	//Prefix("/wx/new/new").
			//	Build(),
			// 基于分支创建新的proto文件，submodule 切换分支生成stub代码， 同理client使用连调切换同一个分支
			// 维护makefile 使用protoc + go build 统一处理
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

	//v1.RegisterWxBaseHTTPServer()

	pbWxBase.RegisterOfficialAccountHTTPServer(srv, oa)
	pbWxBase.RegisterMiniProgramHTTPServer(srv, mini)
	return srv
}

func CustomerResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	fmt.Println("r.URL.Path:", r.URL.Path)
	if v == nil {
		return nil
	}

	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		nethttp.Redirect(w, r, url, code)
		return nil
	}

	if r.URL.Path == "/wx/new/test" && r.Method == "GET" {
		writeContextType(w, plainContentType)
		data := v.(*pbWxBase.WxCallbackReply)
		_, err := w.Write([]byte(data.Echostr))
		if err != nil {
			return err
		}
	} else {
		codec, _ := http.CodecForRequest(r, "Accept")
		data, err := codec.Marshal(v)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", ContentType(codec.Name()))
		_, err = w.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

var plainContentType = []string{"text/plain; charset=utf-8"}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

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
