package chat

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received Message from client: %s", message.Body)
	return &Message{Body: "Hello from server"}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
