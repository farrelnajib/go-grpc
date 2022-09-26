package accessor

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
)

type GetByIDsInput struct {
	IDs []string
}

type GetByIDsOutput struct {
	Products []product.ProductORM
}

func (ths *accessor) GetByIDs(ctx context.Context, input GetByIDsInput) (*GetByIDsOutput, error) {
	var res GetByIDsOutput

	err := ths.db.Preload("ProductVariant").Where("products.id IN ?", input.IDs).Find(&res.Products).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
