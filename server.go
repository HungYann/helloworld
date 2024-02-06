package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	// 导入grpc包
	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) mustEmbedUnimplementedSearchServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *SearchService) Search(ctx context.Context, r *SearchRequest) (*SearchResponse, error) {
	fmt.Println("开始执行逻辑程序")

	defer fmt.Println("结束执行逻辑程序")
	return &SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9002"

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("开始执行拦截器程序")
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Unknown, "")
	}

	referer := md.Get("Accept-Language")
	fmt.Println(referer[0])

	defer fmt.Println("结束执行拦截器程序")
	return handler(ctx, req)

}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor))
	RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
