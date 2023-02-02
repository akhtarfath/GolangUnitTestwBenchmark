package repository

import (
	"github.com/akhtarfath/GolangUnitTestwBenchmark/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct { // <-- this is the struct that will be mocked
	Mock mock.Mock // <-- this is the struct from github.com/stretchr/testify/mock, use to mock function
}

// FindById is a function to find category by id, return nil if not found and return entity.Category if found
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) // <-- this line is important to mock the function FindById in CategoryRepositoryMock
	// .Get is a function from mock.Arguments struct use to get the first argument
	if arguments.Get(0) == nil { // 0 is the index of the first argument, the first argument is the return value of FindById
		return nil // return nil if the first argument is nil
	}
	category := arguments.Get(0).(entity.Category) // FindById the first argument and convert it to entity.Category
	// return the pointer of category because FindById return pointer of entity.Category
	return &category
}
