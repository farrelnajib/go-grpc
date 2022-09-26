package accessor

import (
	"context"
	"fmt"
	"github.com/farrelnajib/go-rpc/product"
)

type ListInput struct {
	FilterProductInput *product.FilterProductSpec
	SortCriteriaInput  []*product.SortCriteriaSpec
	PageInput          *PageInput
}

type ListOutput struct {
	ID     []string
	DbPage DbPage
}

func (ths *accessor) List(ctx context.Context, input ListInput) (*ListOutput, error) {
	var res ListOutput

	baseQuery := ths.db.Table("products").
		Distinct("products.id").
		Joins("join product_variants on products.id = product_variants.product_id")

	if minPrice := input.FilterProductInput.GetMinimumPrice(); minPrice > 0 {
		baseQuery = baseQuery.Where("product_variants.price >= ?", minPrice)
	}

	if maxPrice := input.FilterProductInput.GetMaximumPrice(); maxPrice > 0 {
		baseQuery = baseQuery.Where("product_variants.price <= ? ", maxPrice)
	}

	if searchText := input.FilterProductInput.GetSearchText(); searchText != "" {
		baseQuery = baseQuery.Where("products.name ILIKE ? OR products.description ILIKE ?", fmt.Sprintf("%%%v%%", searchText), fmt.Sprintf("%%%v%%", searchText))
	}

	baseQuery = baseQuery.Group("products.id")

	var total int64
	var limit = 10
	var offset = 0
	if page := input.PageInput; page != nil {
		limit = int(page.Limit)
		offset = int(page.Offset)
	}

	aggregatedQuery := baseQuery.Offset(offset).Limit(limit)

	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = aggregatedQuery.Scan(&res.ID).Error
	if err != nil {
		return nil, err
	}

	res.DbPage = DbPage{
		Offset:       int32(offset),
		Limit:        int32(limit),
		TotalRecords: total,
	}

	return &res, nil
}
