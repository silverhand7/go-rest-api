package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-rest-api/helpers"
	"github.com/go-rest-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into categories(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helpers.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helpers.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT * from categories where id = ?"
	row, err := tx.QueryContext(ctx, SQL, categoryId)
	helpers.PanicIfError(err)
	defer row.Close()

	category := domain.Category{}
	if row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT * FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
