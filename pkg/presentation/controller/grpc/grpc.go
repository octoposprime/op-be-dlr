package presentation

import (
	"net"

	pp_command "github.com/octoposprime/op-be-dlr/internal/application/presentation/port/command"
	pp_query "github.com/octoposprime/op-be-dlr/internal/application/presentation/port/query"
	pb_dlr "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	pb_error "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
	tgrpc "github.com/octoposprime/op-be-shared/tool/grpc"
	"google.golang.org/grpc"
)

// Grpc is the gRPC API for the application
type Grpc struct {
	pb_error.UnimplementedErorrSvcServer
	pb_dlr.UnimplementedDlrSvcServer
	queryHandler   pp_query.QueryPort
	commandHandler pp_command.CommandPort
}

// NewGrpc creates a new instance of Grpc
func NewGrpc(qh pp_query.QueryPort, ch pp_command.CommandPort) *Grpc {
	api := &Grpc{
		queryHandler:   qh,
		commandHandler: ch,
	}
	return api
}

// Serve starts the API server
func (a *Grpc) Serve(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(tgrpc.Interceptor),
	)
	pb_error.RegisterErorrSvcServer(s, a)
	pb_dlr.RegisterDlrSvcServer(s, a)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
