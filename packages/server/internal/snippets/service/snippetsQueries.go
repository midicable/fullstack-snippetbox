package service

import "fullstack-snippetbox-backend/internal/snippets/repository"

type SnippetsService struct {
	repo *repository.SnippetsRepository
}

func NewSnippetsService(snippetsRepo *repository.SnippetsRepository) *SnippetsService {
	return &SnippetsService{repo: snippetsRepo}
}

func (ss *SnippetsService) GetSnippet() string {
	return "STRING FROM SS: " + ss.repo.GetSnippet()  
}