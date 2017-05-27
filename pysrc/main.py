

import srv_pb2
import srv_pb2_grpc

import grpc

def main():
    channel = grpc.insecure_channel('localhost:50051')
    stub = srv_pb2_grpc.GreeterStub(channel)
    response = stub.SayHello(srv_pb2.HelloRequest(name='you'))
    counter = 1
    try:
        for i in response:
            if counter % 1000 == 0:
                print("got %s messages" % counter)
            counter += 1
    except:
        print("failed after %s messages" % counter)
        raise
    print("Done: %s messages" % counter)

if __name__ == "__main__":
    main()
