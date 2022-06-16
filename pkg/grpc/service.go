package grpc

import (
	"fmt"
	"github.com/petar-arandjic/virtual_kelner.proto/gen/go/business/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Service struct {
	business.UnimplementedBusinessServer
	//appService *internal.AppService
}

// Start grpc server
func Start(port int) {
	s := Service{}

	// Init internal
	//a := internal.NewApp()
	//s.appService = a

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	business.RegisterBusinessServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc: %v", err)
	}
}
