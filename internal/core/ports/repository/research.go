package repository

import "moulaybdl/researchCollab/researchSevice/internal/core/domain"


type ResearchPaperRepository interface {
	// GetAllResearchPapers retrieves all research papers.
	GetAllResearchPapers() ([]domain.ResearchPaper, error)
	// GetResearchPaperByID retrieves a research paper by its ID.
	GetResearchPaperByID(id string) (domain.ResearchPaper, error)
	// GetReasearchPapersByCategory retrieves research papers by category ID.
	GetReasearchPapersByCategory(categoryID string) ([]domain.ResearchPaper, error)
	// CreateResearchPaper creates a new research paper.
	CreateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error)
	// UpdateResearchPaper updates an existing research paper.
	UpdateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error)
	// DeleteResearchPaper deletes a research paper by its ID.
	DeleteResearchPaper(id string) error
}