CREATE TABLE Researcher (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE ResearchPaper (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  short_description TEXT,
  paper_link TEXT NOT NULL,
  researcherID INTEGER REFERENCES Researcher(id) ON DELETE CASCADE
);


CREATE TABLE Category (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL ,
  description TEXT 
);

CREATE TABLE ResearchCategory (
  researchPaperID INTEGER REFERENCES ResearchPaper(id) ON DELETE CASCADE,
  categoryID INTEGER REFERENCES Category(id) ON DELETE CASCADE,
  PRIMARY KEY (researchPaperID, categoryID)
);