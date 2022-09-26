package impl

import (
	"context"
	"fmt"
	"github.com/farrelnajib/go-rpc/product"
)

func (ths *ProductService) UpsertProduct(ctx context.Context, req *product.UpsertProductReq) (*product.UpsertProductRes, error) {
	fmt.Println("Masuk sini")
	result, err := ths.Accessor.Upsert(ctx, req)
	if err != nil {
		return nil, err
	}

	res := product.UpsertProductRes{Uuid: result.Id}
	return &res, nil
}
