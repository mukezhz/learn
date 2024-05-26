package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/mukezhz/learn/golang/grpc/todo/gen/go/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	PORT = ":50051"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	requiredHeader := md.Get("x-required-header")
	aliasHeader := md.Get("another-header")
	nothingHeader := md.Get("x-nothing")
	if len(requiredHeader) == 0 {
		return nil, fmt.Errorf("missing required header")
	}
	log.Printf("Received Headers: %v __ %v __ %v\n", requiredHeader, aliasHeader, nothingHeader)
	todo := &pb.CreateTodoResponse{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	timestampFormat := time.StampNano
	header := metadata.New(
		map[string]string{
			"location":  "Nepal",
			"timestamp": time.Now().Format(timestampFormat),
		})
	grpc.SendHeader(ctx, header)

	return todo, nil
}

var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")

func HeaderMatcher(key string) (string, bool) {
	switch key {
	case "X-Required-Header":
		return key, true
	case "X-Alias-Header":
		return "another-header", true
	case "X-Nothing":
		return key, false
	default:
		return key, true
	}
}

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json+pretty", &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				Indent:    "  ",
				Multiline: true, // Optional, implied by presence of "Indent".
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithIncomingHeaderMatcher(HeaderMatcher),
	)
	err := pb.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}
	log.Println("HTTP server listening on port 8081")
	return http.ListenAndServe(":8081", mux)
}

func runGRPC() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func main() {
	flag.Parse() // Move flag.Parse() here to ensure it's done at the start.

	go runGRPC()
	go func() {
		if err := runHTTP(); err != nil {
			grpclog.Fatal(err)
		}
	}()
	select {} // Block main from exiting
}
