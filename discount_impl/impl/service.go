package impl

import (
	"github.com/farrelnajib/go-rpc/discount"
	"github.com/farrelnajib/go-rpc/discount_impl/accessor"
	"github.com/farrelnajib/go-rpc/product"
)

type DiscountService struct {
	Accessor       accessor.Accessor
	ProductService product.ProductServiceClient
}

func NewDiscountService(
	accessor accessor.Accessor,
	productService product.ProductServiceClient,
) discount.DiscountServiceServer {
	return &DiscountService{
		Accessor:       accessor,
		ProductService: productService,
	}
}
