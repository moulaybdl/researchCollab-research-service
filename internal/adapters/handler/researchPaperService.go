package handler

import (
	"context"
	pb "moulaybdl/researchCollab/researchSevice/internal/core/ports/protobufs/protobufs"
	"moulaybdl/researchCollab/researchSevice/internal/core/ports/services"
	"strconv"

	"google.golang.org/protobuf/types/known/emptypb"
)



type ResearchPaperServer struct {
	pb.UnimplementedResearchPaperServer

	// services:
	rpService services.ResearchPaperServiceInterface
}

func NewResearchPaperServer(rps services.ResearchPaperServiceInterface) *ResearchPaperServer {
	return &ResearchPaperServer{
		rpService: rps,
	}
}


func (s *ResearchPaperServer) GetAllResearchPapers(ctx context.Context, empty *emptypb.Empty) (*pb.ListofPapers, error) {

	// for test: this will call getResearch paper by ID:
	rps, err := s.rpService.GetAllResearchPapers()
	if err != nil {
		return nil, err
	}

	var papers []*pb.Paper
	for _, rp  := range rps {
		// conver id to int:
		int_id, err := strconv.Atoi(rp.ID)
		if err != nil {
			return nil, err
		}

		int_rt_id, err := strconv.Atoi(rp.ResearchTeam)
		if err != nil {
			return nil, err
		}

		publish_year_int, err := strconv.Atoi(rp.PublishedAt)
		if err != nil {
			return nil, err
		}

		paper := pb.Paper {
			Id: int32(int_id),
			Title: rp.Title,
			ShortDescription: rp.ShortDescription,
			PaperLink: rp.PublishedAt,
			ResearcherID: int32(int_rt_id),
			PublishYear: int32(publish_year_int),
		}

		papers = append(papers, &paper)
	}


  return &pb.ListofPapers{Papers: papers}, nil
}