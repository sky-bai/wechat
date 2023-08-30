package server

import (
	"cdgitlib.spreadwin.cn/core/common/pb/pbWxBase"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"wx-base/internal/conf"
	"wx-base/internal/service/mini_program_service"
	"wx-base/internal/service/official_account_service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, oa *official_account_service.OfficialAccountService, mini *mini_program_service.MiniProgramService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pbWxBase.RegisterMiniProgramServer(srv, mini)
	pbWxBase.RegisterOfficialAccountServer(srv, oa)
	return srv
}
