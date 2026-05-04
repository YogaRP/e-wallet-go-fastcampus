package cmd

import (
	"ewallet-fastcampus/helpers"
	"log"
	"net"

	pb "ewallet-fastcampus/cmd/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	dependency := dependencyInject()

	s := grpc.NewServer()

	//list method
	pb.RegisterTokenValidationServer(s, dependency.TokenValidationApi)

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))

	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	logrus.Info("start listening grpc on port: " + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port ", err)
	}
}
