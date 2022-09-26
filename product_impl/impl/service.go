package impl

import (
	"github.com/farrelnajib/go-rpc/product"
	"github.com/farrelnajib/go-rpc/product_impl/accessor"
)

type ProductService struct {
	Accessor accessor.Accessor
}

func NewProductService(
	accessor accessor.Accessor,
) product.ProductServiceServer {
	return &ProductService{Accessor: accessor}
}
