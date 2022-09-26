package impl

import (
	"context"
	"fmt"
	"github.com/farrelnajib/go-rpc/discount"
	"github.com/farrelnajib/go-rpc/discount_impl/accessor"
	"github.com/farrelnajib/go-rpc/product"
	"time"
)

func (ths *DiscountService) GetDiscountedProductIDs(ctx context.Context, req *discount.GetDiscountedProductIDsReq) (*discount.GetDiscountedProductIDsRes, error) {
	input := accessor.GetByIDsInput{IDs: []string{req.DiscountId}}

	res, err := ths.Accessor.GetByIDs(ctx, input)
	if err != nil {
		return nil, err
	}

	productIDs := []string{}
	for _, disc := range res.Discounts {
		for _, row := range disc.Rows {
			sku := row.SKU
			req := product.GetProductIDsBySKUsReq{SKUs: []string{sku}}
			now := time.Now()
			res, err := ths.ProductService.GetProductIDsBySKUs(ctx, &req)
			fmt.Println("Time delay", time.Since(now))
			if err != nil {
				return nil, err
			}

			ids := res.GetIDs()
			for _, id := range ids {
				productIDs = append(productIDs, id)
			}
		}
	}

	return &discount.GetDiscountedProductIDsRes{ProductIds: productIDs}, nil
}
