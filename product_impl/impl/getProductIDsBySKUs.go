package impl

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
)

func (ths *ProductService) GetProductIDsBySKUs(ctx context.Context, req *product.GetProductIDsBySKUsReq) (*product.GetProductIDsBySKUsRes, error) {
	input := accessor.GetProductIDsBySKUsInput{SKUs: req.GetSKUs()}

	output, err := ths.Accessor.GetProductIDsBySKUs(ctx, input)
	if err != nil {
		return nil, err
	}

	return &product.GetProductIDsBySKUsRes{IDs: output.IDs}, nil
}
