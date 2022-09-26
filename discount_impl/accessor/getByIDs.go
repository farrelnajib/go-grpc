package accessor

import (
	"context"
	"github.com/farrelnajib/go-rpc/discount"
)

type GetByIDsInput struct {
	IDs []string
}

type GetByIDsOutput struct {
	Discounts []discount.DiscountORM
}

func (ths *accessor) GetByIDs(ctx context.Context, input GetByIDsInput) (*GetByIDsOutput, error) {
	var res GetByIDsOutput

	err := ths.db.Preload("Rows").Where("discounts.id IN ?", input.IDs).Find(&res.Discounts).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
