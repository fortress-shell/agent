package rpc

import (
  "context"
)

type AgentServer struct {}

func (s *AgentServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
  return &HelloReply{Message: "Hello " + in.Name}, nil
}
