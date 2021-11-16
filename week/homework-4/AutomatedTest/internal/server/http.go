package server

import (
	v1 "HomeworkFour/api/helloworld/v1"
	pb "HomeworkFour/api/trigger/v1"
	"HomeworkFour/internal/conf"
	"HomeworkFour/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, trigger *service.TriggerService,logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	openAPIhandler := openapiv2.NewHandler()
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	pb.RegisterTriggerHTTPServer(srv,trigger)
	return srv
}
