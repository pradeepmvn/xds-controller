package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Pallinder/sillyname-go"
	personpb "github.com/pradeepmvn/xds-controller/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Name server reponds with random person name
type server struct {
	personpb.UnimplementedPersonServer
}

var (
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	serverID   int32
)

func main() {
	// get Env variables
	serverPort := getEnv("SERVER_PORT", "5432")
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverID = randNum(3)
	s := grpc.NewServer()
	personpb.RegisterPersonServer(s, &server{})
	log.Println(" Server Started !!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetDetails(ctx context.Context, in *personpb.PersonRequest) (*personpb.PersonResponse, error) {
	log.Println("Served Request from Server: ", serverID)
	// randomly send unavailable
	if time.Now().Unix()%3 == 0 {
		return nil, status.Errorf(codes.Unavailable, "Triggering Unavailable for retry check")
	}
	return &personpb.PersonResponse{Name: sillyname.GenerateStupidName(), Id: serverID}, nil
}

func getEnv(key, defaultVal string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return v
}

// 9 for ssn
func randNum(length int) int32 {
	var b string
	for i := 0; i < length; i++ {
		b = b + strconv.Itoa(seededRand.Intn(length))
	}
	r, _ := strconv.Atoi(b)
	return int32(r)
}
