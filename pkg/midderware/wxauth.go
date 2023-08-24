package midderware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func WxAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				fmt.Println("start", tr.Operation())

				defer func() {
					// Do something on exiting
					fmt.Println("exit")
				}()
			}
			return handler(ctx, req)
		}
	}
}

// 1.获取到该小程序或者是公众号的配置信息
