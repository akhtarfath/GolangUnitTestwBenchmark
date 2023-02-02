package service

import (
	"github.com/akhtarfath/GolangUnitTestwBenchmark/entity"
	"github.com/akhtarfath/GolangUnitTestwBenchmark/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock nya
	// Mock.On is a function from mock.Mock struct use to mock function FindById in CategoryRepositoryMock
	// feel free to change the function name to anything you want in On("FindById", "1").Return(nil)
	categoryRepository.Mock.On("FindById", "1").Return(nil) // <-- this line is important to mock the function FindById in CategoryRepositoryMock
	// categoryService.FindById("1") will call CategoryRepositoryMock.FindById("1")
	category, err := categoryService.FindById("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	// insert mock data to Category Entity
	category := entity.Category{
		Id:   "2",
		Name: "Handphone",
	}
	// program mock nya
	// Mock.On is a function from mock.Mock struct use to mock function FindById in CategoryRepositoryMock
	// feel free to change the function name to anything you want in On("FindById", "2").Return(category)
	categoryRepository.Mock.On("FindById", "2").Return(category) // <-- this line is important to mock the function FindById in CategoryRepositoryMock
	// categoryService.FindById("2") will call CategoryRepositoryMock.FindById("2")
	result, err := categoryService.FindById("2")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equalf(t, "2", result.Id, "Id should be 2")
	assert.Equalf(t, "Handphone", result.Name, "Name should be Handphone")
}
