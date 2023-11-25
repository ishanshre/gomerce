package dbrepo

import (
	"context"

	"github.com/ishanshre/gomerce/internals/model"
)

func (r *postgresDBRepo) CreateCategory(name string) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()
	category := &model.Category{}
	query := `INSERT INTO categories (name) VALUES $1 RETURNING id, name;`
	if err := r.DB.GetDB().QueryRowContext(ctx, query, name).Scan(
		&category.Id,
		&category.Name,
	); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *postgresDBRepo) GetCategories() ([]*model.Category, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `SELECT * FROM categories;`
	rows, err := r.DB.GetDB().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	categories := []*model.Category{}
	for rows.Next() {
		category := &model.Category{}
		if err := rows.Scan(
			&category.Id,
			&category.Name,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *postgresDBRepo) GetCategory(id int) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `SELECT * FROM categories WHERE id=$1;`
	category := &model.Category{}
	if err := r.DB.GetDB().QueryRowContext(ctx, query, id).Scan(
		&category.Id,
		&category.Name,
	); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *postgresDBRepo) DeleteCategory(id int) error {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `DELETE FROM categories WHERE id=$1;`
	_, err := r.DB.GetDB().ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresDBRepo) UpdateCategory(update_data *model.Category) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `
		UPDATE categories
		SET name = $2
		WHERE id=$1
		RETURNING id, name;
		`
	updated_category := &model.Category{}
	if err := r.DB.GetDB().QueryRowContext(
		ctx,
		query,
		update_data.Id,
		update_data.Name,
	).Scan(
		&updated_category.Id,
		&updated_category.Name,
	); err != nil {
		return nil, err
	}
	return updated_category, nil
}
