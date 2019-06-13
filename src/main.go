package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"./api/proto/pb_src"
	"./api/proto/v1"
	"./models"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC | log.Lshortfile)

	models.InitDB()

	log.Println("Server is starting on port", os.Getenv("MAIN_PORT"))

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", os.Getenv("MAIN_PORT")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := v1.Server{}

	grpcServer := grpc.NewServer()

	pb_src.RegisterLoggerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
