package impl

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
)

func (ths *ProductService) List(ctx context.Context, req *product.ListProductReq) (*product.ListProductRes, error) {
	output, err := ths.Accessor.List(ctx, accessor.ListInput{
		FilterProductInput: req.GetFilterProductSpec(),
		PageInput:          accessor.ConstructPageInputFromPageSpec(req.GetPageSpec()),
		SortCriteriaInput:  req.GetSortCriteriaSpec(),
	})
	if err != nil {
		return nil, err
	}

	return &product.ListProductRes{
		ProductIDs: output.ID,
		PageInfo:   accessor.ConstructPageInfoFromDbPage(output.DbPage),
	}, nil
}
