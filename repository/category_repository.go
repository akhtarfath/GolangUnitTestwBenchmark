package repository

import "github.com/akhtarfath/GolangUnitTestwBenchmark/entity"

// CategoryRepository is a repository for category
type CategoryRepository interface { // <-- this is the interface that will be mocked
	FindById(id string) *entity.Category // <-- this is the function that will be mocked
}
