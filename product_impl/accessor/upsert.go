package accessor

import (
	"context"
	"errors"
	"github.com/farrelnajib/go-rpc/product"
	"github.com/google/uuid"
)

func (ths *accessor) Upsert(ctx context.Context, input *product.UpsertProductReq) (*product.ProductORM, error) {
	productORM, err := constructUpsertInput(input)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = ths.db.Create(&productORM).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	variants := constructVariant(input, productORM)
	err = ths.db.Create(&variants).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return productORM, nil
}

func constructUpsertInput(input *product.UpsertProductReq) (*product.ProductORM, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &product.ProductORM{
		Description: input.GetDescription(),
		Id:          id.String(),
		Name:        input.Name,
		Slug:        input.GetSlug(),
	}, nil
}

func constructVariant(input *product.UpsertProductReq, prd *product.ProductORM) []*product.ProductVariantORM {
	var variants []*product.ProductVariantORM
	for _, variant := range input.GetProductVariants() {
		variantORM := product.ProductVariantORM{
			ProductId: &prd.Id,
			Price:     variant.GetPrice(),
			SKU:       variant.GetSKU(),
		}

		variants = append(variants, &variantORM)
	}

	return variants
}
