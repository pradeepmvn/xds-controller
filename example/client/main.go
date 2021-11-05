package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	personpb "github.com/pradeepmvn/xds-controller/example/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
)

func main() {
	// get Env variables
	server := getEnv("SERVER_URL", "localhost:5432")
	sleepTime, _ := strconv.Atoi(getEnv("SLEEP_TIME", "1"))
	runTime, _ := strconv.Atoi(getEnv("TOTAL_RUN_TIME", "60"))
	// start connection
	log.Printf("Establishing connection to : %s", server)
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Panic("Unable to connect to server", err)
	}
	defer conn.Close()
	client := personpb.NewPersonClient(conn)
	now := time.Now()
	log.Printf("Starting the Process")
	for {
		r, err := client.GetDetails(context.Background(), &personpb.PersonRequest{})
		if err != nil {
			// Will crash container when Unavailable is sent from server
			log.Fatalf("could not get details: %v", err)
		}
		log.Printf("Got Name: %s from Server: %d", r.GetName(), r.GetId())
		time.Sleep(time.Duration(sleepTime) * time.Second)
		if time.Now().Sub(now) >= time.Duration(time.Duration(runTime)*time.Minute) {
			break
		}
	}
	log.Printf("Completed cycle of %v sec. Killing the server", runTime)
}

func getEnv(key, defaultVal string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return v
}
