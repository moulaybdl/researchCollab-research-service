package handler

import (
	"context"
	"moulaybdl/researchCollab/researchSevice/internal/core/domain"
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

func (s *ResearchPaperServer) GetResearchPaperByCategory(ctx context.Context, category_req *pb.GetResearchPaperByCategoryRequest) (*pb.ListofPapers, error) {
	rps, err := s.rpService.GetReasearchPapersByCategory(category_req.Category)
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

func (s *ResearchPaperServer) GetResearchPaperByID(ctx context.Context, id_req *pb.GetResearchPaperByIDRequest) (*pb.Paper, error) {
	rp, err := s.rpService.GetResearchPaperByID(id_req.PaperId)
	if err != nil {
		return nil, err
	}

	var paper pb.Paper

	paper_id_int, err := strconv.Atoi(rp.ID)
	if err != nil {
		return nil, err
	}
	paper.Id = int32(paper_id_int)

	paper.PaperLink = rp.Link

	publish_int, err := strconv.Atoi(rp.PublishedAt)
	if err != nil {
		return nil, err
	}
	paper.PublishYear = int32(publish_int)


	researcher_id_int, err := strconv.Atoi(rp.ResearchTeam)
	if err != nil {
		return nil, err
	}
	paper.ResearcherID = int32(researcher_id_int)



	paper.ShortDescription = rp.ShortDescription
	paper.Title = rp.Title

	return &paper, nil
}

func (s *ResearchPaperServer) CreateResearchPaper(ctx context.Context, rp_req *pb.CreateResearchPaperRequest) (*pb.CreateResearchPaperResponse, error) {
	var paper_input domain.ResearchPaper
	paper_input.Link = rp_req.PaperLink
	paper_input.Description = rp_req.ShortDescription
	paper_input.Title = rp_req.Title
	paper_input.PublishedAt = rp_req.PublishYear
	paper_input.ResearchTeam = rp_req.ResearcherId

	p, err := s.rpService.CreateResearchPaper(paper_input)
	if err != nil {
		return &pb.CreateResearchPaperResponse{Id: "-1"}, err
	}

	var paper_reponse pb.CreateResearchPaperResponse
	paper_reponse.Id = p.ID
	paper_reponse.PaperLink = p.Link
	paper_reponse.PublishYear = p.PublishedAt
	paper_reponse.ResearcherId = p.ResearchTeam
	paper_reponse.ShortDescription = p.ShortDescription
	paper_reponse.Title = p.Title

	return &paper_reponse, nil
}

func (s *ResearchPaperServer) DeleteResearchPaper(ctx context.Context, rp_req *pb.DeleteResearchPaperRequest) (*pb.DeleteResearchPaperResponse, error) {
	err := s.rpService.DeleteResearchPaper(rp_req.Id)
	if err != nil {
		return &pb.DeleteResearchPaperResponse{State: "failed"}, err
	}

	return &pb.DeleteResearchPaperResponse{State: "success"}, nil
}