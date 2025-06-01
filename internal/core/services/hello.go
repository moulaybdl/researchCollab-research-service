package services

import (
	"context"
	pb "moulaybdl/researchCollab/researchSevice/internal/core/ports/protobufs/protobufs"
)

type Server struct {
    pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}