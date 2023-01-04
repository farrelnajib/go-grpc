package impl

import (
	"context"
	"fmt"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
	"time"
)

func (ths *ProductService) GetProductIDsBySKUs(ctx context.Context, req *product.GetProductIDsBySKUsReq) (*product.GetProductIDsBySKUsRes, error) {
	now := time.Now()
	input := accessor.GetProductIDsBySKUsInput{SKUs: req.GetSKUs()}

	output, err := ths.Accessor.GetProductIDsBySKUs(ctx, input)
	if err != nil {
		return nil, err
	}

	fmt.Println("Time delay", time.Since(now))
	return &product.GetProductIDsBySKUsRes{IDs: output.IDs}, nil
}
