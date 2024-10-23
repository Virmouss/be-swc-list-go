package main

import (
	controller "be-swc-list/app/controller"
	"be-swc-list/app/db"
	"be-swc-list/app/services"
	"be-swc-list/utils"
	"log"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDB()
	utils.Seed(db)

	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		log.Println("Connected to database")
	}

	StartGRPCServer(9000)
	StartGinServer(9090)
}

func StartGRPCServer(port int) {
	portNumber := strconv.Itoa(port)
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("gRPC fail to listen on port "+portNumber+" : %v", err)
	} else {
		log.Println("gRPC Starting...")

		grpcServer := grpc.NewServer()

		//register service
		swcService := services.NewSensorWeaponCoverageService()
		controller.NewGrpcSWCService(grpcServer, swcService)

		go func() {
			log.Println("gRPC listening to port " + portNumber)
			if err := grpcServer.Serve(lis); err != nil {
				log.Fatalf("could not start grpc server: %v", err)
			}
		}()
		//log.Println("gRPC listening to port " + portNumber)
	}
}

func StartGinServer(port int) {
	portNumber := strconv.Itoa(port)

	server := gin.Default()

	server.Run("localhost:" + portNumber)
}
