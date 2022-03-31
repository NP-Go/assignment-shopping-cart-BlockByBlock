package types

type ItemInfo struct {
	category int
	qty      int
	unitCost float64
}

func (itemInfo *ItemInfo) SetCategory(category int) {
	itemInfo.category = category
}

func (itemInfo *ItemInfo) SetQty(qty int) {
	itemInfo.qty = qty
}

func (itemInfo *ItemInfo) SetUnitCost(unitCost float64) {
	itemInfo.unitCost = unitCost
}
