package repository

import entity "learn-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
