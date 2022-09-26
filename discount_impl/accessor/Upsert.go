package accessor

import (
	"context"
	"errors"
	"github.com/farrelnajib/go-rpc/discount"
	"github.com/google/uuid"
)

func (ths *accessor) Upsert(ctx context.Context, input *discount.UpsertDiscountReq) (*discount.DiscountORM, error) {
	discountORM, err := constructUpsertInput(input)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = ths.db.Create(&discountORM).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	rows := constructDiscountRows(input, discountORM.Id)
	err = ths.db.Create(&rows).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return discountORM, nil
}

func constructUpsertInput(input *discount.UpsertDiscountReq) (*discount.DiscountORM, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	startDate := input.GetStartDate().AsTime()
	endDate := input.GetEndDate().AsTime()

	return &discount.DiscountORM{
		EndDate:   &endDate,
		Id:        id.String(),
		StartDate: &startDate,
		Title:     input.GetTitle(),
	}, nil
}

func constructDiscountRows(input *discount.UpsertDiscountReq, id string) []*discount.RowORM {
	var rows []*discount.RowORM
	for _, row := range input.GetRows() {
		rowORM := discount.RowORM{
			DiscountId: &id,
			SKU:        row.GetSKU(),
			Value:      row.GetValue(),
		}

		rows = append(rows, &rowORM)
	}

	return rows
}