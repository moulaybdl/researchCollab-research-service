package repository

import "moulaybdl/researchCollab/researchSevice/internal/core/domain"


func (r *DB) GetAllResearchPapers() ([]domain.ResearchPaper, error) {
	query := `
	SELECT * FROM researchpaper
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var papers []domain.ResearchPaper
	for rows.Next() {
		var paper domain.ResearchPaper
		err := rows.Scan(&paper.ID, &paper.Title, &paper.ShortDescription, &paper.Link, &paper.ResearchTeam, &paper.PublishedAt)
		if err != nil {
			return nil, err
		}
		papers = append(papers, paper)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return papers, nil
}

func (r *DB) GetResearchPaperByID(id string) (domain.ResearchPaper, error) {
	query := `
	SELECT * FROM researchpaper WHERE id = $1
	`

	var paper domain.ResearchPaper
	err := r.db.QueryRow(query, id).Scan(&paper.ID, &paper.Title, &paper.ShortDescription, &paper.Link, &paper.ResearchTeam, &paper.PublishedAt)
	if err != nil {
		return domain.ResearchPaper{}, err
	}

	return paper, nil
}

func (r *DB) GetReasearchPapersByCategory(categoryID string) ([]domain.ResearchPaper, error) {
	query := `
	SELECT rp.* FROM researchpaper rp
	JOIN researchcategory rpc ON rp.id = rpc.researchpaper_id
	WHERE rpc.category_id = $1
	`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var papers []domain.ResearchPaper
	for rows.Next() {
		var paper domain.ResearchPaper
		err := rows.Scan(&paper.ID, &paper.Title, &paper.ShortDescription, &paper.Link, &paper.ResearchTeam, &paper.PublishedAt)
		if err != nil {
			return nil, err
		}
		papers = append(papers, paper)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return papers, nil
}

func (r *DB) GetResearchPapersByResearcher(researcherID string) ([]domain.ResearchPaper, error) {
	query := `
	SELECT * FROM researchpaper 
	WHERE researcherid = $1
	`

	rows, err := r.db.Query(query, researcherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var papers []domain.ResearchPaper
	for rows.Next() {
		var paper domain.ResearchPaper
		err := rows.Scan(&paper.ID, &paper.Title, &paper.ShortDescription, &paper.Link, &paper.ResearchTeam, &paper.PublishedAt)
		if err != nil {
			return nil, err
		}
		papers = append(papers, paper)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return papers, nil
}

func (r *DB) CreateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error) {
	query := `
	INSERT INTO researchpaper (title, shortdescription, paper_link, researcherid, publish_year)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err := r.db.QueryRow(query, paper.Title, paper.ShortDescription, paper.Link, paper.ResearchTeam, paper.PublishedAt).Scan(&paper.ID)
	if err != nil {
		return domain.ResearchPaper{}, err
	}

	return paper, nil
}

func (r *DB) UpdateResearchPaper(paper domain.ResearchPaper) (domain.ResearchPaper, error) {
	query := `
	UPDATE researchpaper
	SET title = $1, shortdescription = $2, paper_link = $3, researcherid = $4, publish_year = $5
	WHERE id = $6
	RETURNING id
	`

	err := r.db.QueryRow(query, paper.Title, paper.ShortDescription, paper.Link, paper.ResearchTeam, paper.PublishedAt, paper.ID).Scan(&paper.ID)
	if err != nil {
		return domain.ResearchPaper{}, err
	}

	return paper, nil
}

func (r *DB) DeleteResearchPaper(id string) error {
	query := `
	DELETE FROM researchpaper WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

