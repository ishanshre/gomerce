package repository

import "github.com/ishanshre/gomerce/internals/model"

type Repository interface {

	// category interface
	CreateCategory(name string) (*model.Category, error)
	GetCategories() ([]*model.Category, error)
	GetCategory(id int) (*model.Category, error)
	DeleteCategory(id int) error
	UpdateCategory(update_data *model.Category) (*model.Category, error)
}
