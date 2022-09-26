package accessor

import "github.com/farrelnajib/go-rpc/product"

func ConstructPageInputFromPageSpec(pageSpec *product.PageSpec) *PageInput {
	if pageSpec != nil {
		pageNumber := pageSpec.PageNumber
		if pageNumber < 0 {
			pageNumber = 0
		}

		itemPerPage := pageSpec.ItemPerPage
		if itemPerPage < 1 {
			itemPerPage = 10
		}

		return &PageInput{
			Offset: itemPerPage * pageNumber,
			Limit:  itemPerPage,
		}
	}

	return nil
}

func ConstructPageInfoFromDbPage(dbPage DbPage) *product.PageInfo {
	return &product.PageInfo{
		ItemPerPage: dbPage.Limit,
		PageNumber:  dbPage.Offset / dbPage.Limit,
		TotalItems:  int32(dbPage.TotalRecords),
	}
}
