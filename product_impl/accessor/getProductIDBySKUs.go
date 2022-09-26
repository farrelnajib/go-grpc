package accessor

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
)

type GetProductIDsBySKUsInput struct {
	SKUs []string
}

type GetProductIDsBySKUsOutput struct {
	IDs []string
}

func (ths *accessor) GetProductIDsBySKUs(ctx context.Context, input GetProductIDsBySKUsInput) (*GetProductIDsBySKUsOutput, error) {
	var res GetProductIDsBySKUsOutput

	err := ths.db.Model(product.ProductORM{}).
		Select("products.id").
		Joins("JOIN product_variants ON products.id = product_variants.product_id").
		Where("product_variants.SKU IN ?", input.SKUs).
		Scan(&res.IDs).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
