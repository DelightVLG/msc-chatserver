package main

import (
	"context"
	"log"
	"net"

	chatv1 "github.com/DelightVLG/msc-chatserver/pkg/api/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	chatv1.UnimplementedChatServiceServer
}

const (
	port = ":8088"
)

func (s *server) Create(ctx context.Context, req *chatv1.CreateRequest) (*chatv1.CreateResponse, error) {
	id := int64(1)

	return &chatv1.CreateResponse{
		Id: &id,
	}, nil

}

func (s *server) Delete(ctx context.Context, req *chatv1.DeleteRequest) (*chatv1.DeleteResponse, error) {
	return &chatv1.DeleteResponse{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	return &chatv1.SendMessageResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	chatv1.RegisterChatServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
