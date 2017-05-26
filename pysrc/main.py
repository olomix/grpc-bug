

import srv_pb2
import srv_pb2_grpc

import grpc

def main():
    channel = grpc.insecure_channel('localhost:50051')
    stub = srv_pb2_grpc.GreeterStub(channel)
    response = stub.SayHello(srv_pb2.HelloRequest(name='you'))
    counter = 1
    for i in response:
        print i, counter
        counter += 1

if __name__ == "__main__":
    main()
