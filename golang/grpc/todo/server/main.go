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

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	customHeader1 := md.Get("x-custom-header1")
	customHeader2 := md.Get("new-header")
	customHeader3 := md.Get("custom-nothing")
	log.Printf("Received Headers: %v __ %v __ %v\n", customHeader1, customHeader2, customHeader3)
	todo := &pb.CreateTodoResponse{
		Name:        in.GetName(),
		Description: in.GetDescription(),
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
	case "X-Custom-Header1":
		return key, true
	case "X-Custom-Header2":
		return "new-header", true
	case "X-Custom-Nothing":
		return "custom-nothing", false
	default:
		fmt.Println("HeaderMatcher", key)
		return key, false
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
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
		if err := run(); err != nil {
			grpclog.Fatal(err)
		}
	}()
	select {} // Block main from exiting
}
