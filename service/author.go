package service

import (
	"base-gin/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: authorRepo}
}