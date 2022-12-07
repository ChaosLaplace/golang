package controllers

import (
	pb "Heroku/proto"
	"context"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	features []*pb.Feature
}

func (s *routeGuideServer) GetFeature(cxt context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

func GrpcServer(c *gin.Context) {
	fmt.Printf("GrpcServer Listen \n")
	// 生成一個listener
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Printf("GrpcServer net.Listen | err=%v \n", err)
	}
	// server
	fmt.Printf("GrpcServer NewServer \n")
	grpcServer := grpc.NewServer()
	fmt.Printf("GrpcServer RegisterRouteGuideServer \n")
	pb.RegisterRouteGuideServer(grpcServer, dbServer())
	fmt.Printf("GrpcServer Serve=%v", grpcServer.Serve(lis))
}

func dbServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{
				Name: "東京鐵塔",
				Location: &pb.Point{
					Latitude:  353931000,
					Longitude: 139444400,
				},
			},
			{
				Name: "淺草寺",
				Location: &pb.Point{
					Latitude:  357147651,
					Longitude: 139794466,
				},
			},
			{
				Name: "晴空塔",
				Location: &pb.Point{
					Latitude:  357100670,
					Longitude: 139808511,
				},
			},
		},
	}
}
