package repositories

import (
	"context"

	"github.com/cr1m3s/tch_backend/queries"
)

type CategoriesRepositoryInterface interface {
	GetCategoriesWithChildren(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error)
	GetCategoryAndParent(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error)
	GetCategoryByID(ctx context.Context, id int32) (queries.Category, error)
	GetCategoryByName(ctx context.Context, name string) (queries.Category, error)
	GetCategoryParents(ctx context.Context) ([]queries.Category, error)
}

type CategoriesRepository struct {
	q *queries.Queries
}

func NewCategoriesRepository() *CategoriesRepository {
	return &CategoriesRepository{
		q: NewAppRepository(),
	}
}

func (t *CategoriesRepository) GetCategoriesWithChildren(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error) {
	return t.q.GetCategoriesWithChildren(ctx)
}

func (t *CategoriesRepository) GetCategoryAndParent(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error) {
	return t.q.GetCategoryAndParent(ctx, name)
}

func (t *CategoriesRepository) GetCategoryByID(ctx context.Context, id int32) (queries.Category, error) {
	return t.q.GetCategoryByID(ctx, id)
}

func (t *CategoriesRepository) GetCategoryByName(ctx context.Context, name string) (queries.Category, error) {
	return t.q.GetCategoryByName(ctx, name)
}

func (t *CategoriesRepository) GetCategoryParents(ctx context.Context) ([]queries.Category, error) {
	return t.q.GetCategoryParents(ctx)
}
