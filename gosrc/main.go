package main

import(
	pb "protobuf"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"time"
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
		time.Sleep(50*time.Microsecond)
		log.Printf("ok %v", i)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
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
