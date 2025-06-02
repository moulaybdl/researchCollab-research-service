package handler

import (
	"context"
	pb "moulaybdl/researchCollab/researchSevice/internal/core/ports/protobufs/protobufs"
)



type ResearchPaperServer struct {
	pb.UnimplementedGreeterServer

	// services:
	


}


func (s *ResearchPaperServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}