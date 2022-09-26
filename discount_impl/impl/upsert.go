package impl

import (
	"context"
	"github.com/farrelnajib/go-rpc/discount"
)

func (ths *DiscountService) UpsertDiscount(ctx context.Context, req *discount.UpsertDiscountReq) (*discount.UpsertDiscountRes, error) {
	result, err := ths.Accessor.Upsert(ctx, req)
	if err != nil {
		return nil, err
	}

	res := discount.UpsertDiscountRes{Uuid: result.Id}
	return &res, nil
}
