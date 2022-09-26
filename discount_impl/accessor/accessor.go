package accessor

import (
	"context"
	"github.com/farrelnajib/go-rpc/discount"
	"gorm.io/gorm"
)

type Accessor interface {
	Upsert(ctx context.Context, input *discount.UpsertDiscountReq) (*discount.DiscountORM, error)
	GetByIDs(ctx context.Context, input GetByIDsInput) (*GetByIDsOutput, error)
}

type accessor struct {
	db *gorm.DB
}

func NewAccessor(db *gorm.DB) Accessor {
	return &accessor{db: db}
}
