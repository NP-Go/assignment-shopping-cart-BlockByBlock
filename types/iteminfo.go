package types

type ItemInfo struct {
	Category int
	Qty      int
	UnitCost float64
}

func (itemInfo *ItemInfo) SetCategory(category int) {
	itemInfo.Category = category
}

func (itemInfo *ItemInfo) SetQty(qty int) {
	itemInfo.Qty = qty
}

func (itemInfo *ItemInfo) SetUnitCost(unitCost float64) {
	itemInfo.UnitCost = unitCost
}
