package impl

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
)

func (ths *ProductService) GetProductByIDs(ctx context.Context, req *product.GetProductByIDsReq) (*product.GetProductByIDsRes, error) {
	input := accessor.GetByIDsInput{IDs: req.GetProductIDs()}
	products, err := ths.Accessor.GetByIDs(ctx, input)
	if err != nil {
		return nil, err
	}

	var res []*product.Product

	for _, prod := range products.Products {
		prodPB, err := prod.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		res = append(res, &prodPB)
	}

	return &product.GetProductByIDsRes{Products: res}, nil
}
