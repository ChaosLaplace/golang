package controllers

import (
	pb "Heroku/proto"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GrpcClient(c *gin.Context) {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("GrpcClient grpc.Dial | err=%v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	feature, err := client.GetFeature(
		context.Background(),
		&pb.Point{
			Latitude:  353931000,
			Longitude: 139444400,
		},
	)
	if err != nil {
		fmt.Printf("GrpcClient client.GetFeature | err=%v", err)
	}

	fmt.Printf("GrpcClient | feature=%v", feature)
}
