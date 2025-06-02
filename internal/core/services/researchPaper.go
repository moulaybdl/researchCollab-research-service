package services

import (
	"moulaybdl/researchCollab/researchSevice/internal/core/domain"
	"moulaybdl/researchCollab/researchSevice/internal/core/ports/repository"
)


type ResearchPaperService struct {
	// include repos here
	researchPaperRepo repository.ResearchPaperRepository
}

func NewResearchPaperService(researchPaperRepo repository.ResearchPaperRepository) *ResearchPaperService{
	return &ResearchPaperService{
		researchPaperRepo: researchPaperRepo,
	}
}


func (s *ResearchPaperService) GetAllResearchPapers() ([]domain.ResearchPaper, error) {
	return s.researchPaperRepo.GetAllResearchPapers()
}


func (s *ResearchPaperService) GetResearchPaperByID(id string) (domain.ResearchPaper, error) {
	return s.researchPaperRepo.GetResearchPaperByID(id)
}

func (s *ResearchPaperService) GetReasearchPapersByCategory(categoryID string) ([]domain.ResearchPaper, error) {
	return s.researchPaperRepo.GetReasearchPapersByCategory(categoryID)
}

func (s *ResearchPaperService) GetResearchPapersByResearcher(researcherID string) ([]domain.ResearchPaper, error) {
	return s.researchPaperRepo.GetResearchPapersByResearcher(researcherID)
}

func (s *ResearchPaperService) CreateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error) {
	return s.researchPaperRepo.CreateResearchPaper(paper)
}

func (s *ResearchPaperService) UpdateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error) {
	return s.researchPaperRepo.UpdateResearchPaper(paper)
}

func (s *ResearchPaperService) DeleteResearchPaper(id string) error {
	return s.researchPaperRepo.DeleteResearchPaper(id)
}



