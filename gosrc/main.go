package main

import(
	pb "protobuf"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"net/http"
)

const (
	grpcListenAddr = ":50051"
	httpListenAddr = ":50052"
)

type server struct {}

func (s *server) SayHello(req *pb.HelloRequest, srv pb.Greeter_SayHelloServer) error {
	log.Printf("Request: %v", req)
	for i := 0; i < 50000; i++ {
		err := srv.Send(&pb.HelloReply{fmt.Sprintf("i=%v", i)})
		if err != nil {
			log.Printf("req %v failed", i)
			return err
		}
		if i % 1000 == 0 {
			log.Printf("ok %v", i)
		}
	}
	return nil
}

func runDebugHttp() {
	go func() {
		log.Fatal(http.ListenAndServe(httpListenAddr, nil))
	}()
}

func main() {
	runDebugHttp()

	lis, err := net.Listen("tcp", grpcListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
