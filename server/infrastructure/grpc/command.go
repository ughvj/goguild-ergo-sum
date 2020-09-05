package grpc

import (
	"net"
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/ughvj/goguild-ergo-sum/pb"
	"github.com/ughvj/goguild-ergo-sum/service"
)

func LaunchCommandServer() error {
	port, err := net.Listen("tcp", ":" + os.Getenv("WORKING_PORT"))
	if err != nil {
		return err
	}
	defer port.Close()
	
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authguard),
		),
	)
	commandService := &service.CommandService{}

	pb.RegisterCommandServer(
		server,
		commandService,
	)
	reflection.Register(server)
	server.Serve(port)
	return nil
}

func authguard(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if token != "hoge" {
		return nil, grpc.Errorf(codes.PermissionDenied, "You are not member of our guild !")
	}
	return context.WithValue(ctx, "result", "ok"), nil
}