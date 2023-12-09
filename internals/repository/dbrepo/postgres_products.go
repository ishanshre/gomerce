package dbrepo

import (
	"context"
	"fmt"

	"github.com/ishanshre/gomerce/internals/model"
)

func (r *postgresDBRepo) GetProducts() ([]*model.Product, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()
	products := []*model.Product{}
	query := `SELECT * FROM products`
	rows, err := r.DB.GetDB().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := model.Product{}
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Brand,
			&product.Sku,
			&product.InStock,
			&product.Image,
			&product.Price,
			&product.DiscountedPrice,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, err
}

func (r *postgresDBRepo) GetProduct(id int) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `SELECT * FROM products WHERE id=$1`
	product := model.Product{}
	if err := r.DB.GetDB().QueryRowContext(ctx, query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Brand,
		&product.Sku,
		&product.InStock,
		&product.Image,
		&product.Price,
		&product.DiscountedPrice,
		&product.CategoryId,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *postgresDBRepo) CreateProduct(create *model.ProductNoID) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()
	query := `
		INSERT INTO products (name, description, brand, sku, in_stock, image, price, discounted_price, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING (id, name, description, brand, sku, in_stock, image, price, discounted_price, category_id, created_at, updated_at)
		`
	product := model.Product{}
	if err := r.DB.GetDB().QueryRowContext(
		ctx,
		query,
		create.Name,
		create.Description,
		create.Brand,
		create.Sku,
		create.InStock,
		create.Image,
		create.Price,
		create.DiscountedPrice,
		create.CategoryId,
		create.CreatedAt,
		create.UpdatedAt,
	).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Brand,
		&product.Sku,
		&product.InStock,
		&product.Image,
		&product.Price,
		&product.DiscountedPrice,
		&product.CategoryId,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *postgresDBRepo) DeleteProduct(id int) error {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	query := `DELETE FROM products WHERE id=$1`
	result, err := r.DB.GetDB().ExecContext(ctx, query)
	if err != nil {
		return err
	}
	rows_affected, _ := result.RowsAffected()
	if rows_affected == 0 {
		return fmt.Errorf("failed to delete the product")
	}
	return nil
}

func (r *postgresDBRepo) UpdateProduct(update *model.Product) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, timeout)
	defer cancel()

	updated_product := model.Product{}
	query := `
		UPDATE products
		SET name = $2, description = $3, brand = $4, sku = $5, in_stock = $6, image = $7, price = $8, discounted_price = $9, category_id = $10, created_at = $11, updated_at = $12
		WHERE id=$1
	`
	if err := r.DB.GetDB().QueryRowContext(
		ctx,
		query,
		update.Id,
		update.Name,
		update.Description,
		update.Brand,
		update.Sku,
		update.InStock,
		update.Image,
		update.Price,
		update.DiscountedPrice,
		update.CategoryId,
		update.CreatedAt,
		update.UpdatedAt,
	).Scan(
		&updated_product.Id,
		&updated_product.Name,
		&updated_product.Description,
		&updated_product.Brand,
		&updated_product.Sku,
		&updated_product.InStock,
		&updated_product.Image,
		&updated_product.Price,
		&updated_product.DiscountedPrice,
		&updated_product.CategoryId,
		&updated_product.CreatedAt,
		&updated_product.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &updated_product, nil
}
