package types

type StoreInfo struct {
	categories []string
}

func (storeInfo *StoreInfo) SetCategories(categories []string) {
	storeInfo.categories = categories
}
