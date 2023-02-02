package service

import (
	"errors"
	"github.com/akhtarfath/GolangUnitTestwBenchmark/entity"
	"github.com/akhtarfath/GolangUnitTestwBenchmark/repository"
)

// CategoryService is a service for category
type CategoryService struct {
	Repository repository.CategoryRepository
}

// FindById is a function to find category by id
func (service CategoryService) FindById(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
