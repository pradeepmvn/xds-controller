package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Pallinder/sillyname-go"
	personpb "github.com/pradeepmvn/xds-controller/example/proto"
	"google.golang.org/grpc"
)

// Name server reponds with random person name
type server struct {
	personpb.UnimplementedPersonServer
}

func main() {
	// get Env variables
	serverPort := getEnv("SERVER_PORT", "5432")
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	personpb.RegisterPersonServer(s, &server{})
	log.Println(" Server Started !!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetDetails(ctx context.Context, in *personpb.PersonRequest) (*personpb.PersonResponse, error) {
	log.Println(" Response sent!!")
	return &personpb.PersonResponse{Name: sillyname.GenerateStupidName()}, nil
}

func getEnv(key, defaultVal string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return v
}
