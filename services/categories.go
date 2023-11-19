package services

import (
	"github.com/cr1m3s/tch_backend/di"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/repositories"
	"github.com/gin-gonic/gin"
)

type CategoriesService struct {
	db repositories.CategoriesRepositoryInterface
}

func NewCategoriesService() *CategoriesService {
	return &CategoriesService{
		db: di.NewCategoriesRepository(),
	}
}

func (t *CategoriesService) CatGetAll(ctx *gin.Context) ([]queries.GetCategoriesWithChildrenRow, error) {
	categories, err := t.db.GetCategoriesWithChildren(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (t *CategoriesService) CatGetByID(ctx *gin.Context, id int32) (queries.Category, error) {
	category, err := t.db.GetCategoryByID(ctx, id)

	if err != nil {
		return queries.Category{}, err
	}

	return category, nil
}

func (t *CategoriesService) CatGetByName(ctx *gin.Context, name string) (queries.Category, error) {
	category, err := t.db.GetCategoryByName(ctx, name)

	if err != nil {
		return queries.Category{}, err
	}

	return category, nil
}

func (t *CategoriesService) CatGetFullName(ctx *gin.Context, name string) (queries.GetCategoryAndParentRow, error) {
	categoryName, err := t.db.GetCategoryAndParent(ctx, name)

	if err != nil {
		return queries.GetCategoryAndParentRow{}, err
	}

	return categoryName, nil
}

func (t *CategoriesService) CatGetParets(ctx *gin.Context) ([]queries.Category, error) {
	parents, err := t.db.GetCategoryParents(ctx)

	if err != nil {
		return []queries.Category{}, err
	}

	return parents, nil
}
