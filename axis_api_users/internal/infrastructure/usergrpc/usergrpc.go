package usergrpc

import (
	"net"

	usergrpcpb "github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/usergrpc/proto"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/interface/apigrpc/userserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

//StartGRPC ...
func StartGRPC(hostGRPC string, userServer *userserver.UserServer) error {

	lis, err := net.Listen("tcp", hostGRPC)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	usergrpcpb.RegisterUserInterfaceServer(s, userServer)

	grpclog.Info("Starting GRPC")

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

//NewGRPCLogger ...
func NewGRPCLogger(l grpclog.LoggerV2) {

	grpclog.SetLoggerV2(l)

}
