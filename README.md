
### grpc框架内容

![image.png](https://cdn.nlark.com/yuque/0/2024/png/26815489/1707209576923-883bb85a-7818-4473-afd7-702315aa3c42.png#averageHue=%2319252b&clientId=u904fe800-52b7-4&from=paste&height=500&id=ue2f96511&originHeight=1000&originWidth=2360&originalType=binary&ratio=2&rotation=0&showTitle=false&size=221608&status=done&style=none&taskId=u7d422550-6d84-41e8-a800-51ea5f2c675&title=&width=1180)


```sql
syntax = "proto3";

option go_package="./;helloworld";

service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse) {}
}

message SearchRequest {
  string request = 1;
}

message SearchResponse {
  string response = 1;
}
```

```sql
protoc --go_out=. --go-grpc_out=. helloworld.proto
```

```sql
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
```

### 调用请求

使用postman grpc, 参考postman调用grpc文章https://apifox.com/apiskills/postman-sends-grpc/


![image.png](https://cdn.nlark.com/yuque/0/2024/png/26815489/1707211866233-02224823-c5ea-4458-8f60-f818a8484be1.png#averageHue=%23fdfdfd&clientId=u904fe800-52b7-4&from=paste&height=878&id=u13bb0099&originHeight=1756&originWidth=2244&originalType=binary&ratio=2&rotation=0&showTitle=false&size=171931&status=done&style=none&taskId=u2eed0000-12fb-46d5-895a-e0357bfa0f8&title=&width=1122)
