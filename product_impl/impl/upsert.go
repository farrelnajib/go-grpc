package impl

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
)

func (ths *ProductService) UpsertProduct(ctx context.Context, req *product.UpsertProductReq) (*product.UpsertProductRes, error) {
	result, err := ths.Accessor.Upsert(ctx, req)
	if err != nil {
		return nil, err
	}

	res := product.UpsertProductRes{Uuid: result.Id}
	return &res, nil
}
