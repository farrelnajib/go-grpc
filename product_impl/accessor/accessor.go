package accessor

import (
	"context"
	"github.com/farrelnajib/go-rpc/product"
	"gorm.io/gorm"
)

type Accessor interface {
	Upsert(ctx context.Context, input *product.UpsertProductReq) (*product.ProductORM, error)
	List(ctx context.Context, input ListInput) (*ListOutput, error)
	GetByIDs(ctx context.Context, input GetByIDsInput) (*GetByIDsOutput, error)
	GetProductIDsBySKUs(ctx context.Context, input GetProductIDsBySKUsInput) (*GetProductIDsBySKUsOutput, error)
}

type PageInput struct {
	Offset int32
	Limit  int32
}

type DbPage struct {
	Offset       int32 `json:"offset"`
	Limit        int32 `json:"limit"`
	TotalRecords int64 `json:"total_records"`
}

type accessor struct {
	db *gorm.DB
}

func NewAccessor(db *gorm.DB) Accessor {
	return &accessor{db: db}
}
