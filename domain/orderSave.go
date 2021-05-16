package domain

type OrderSave struct {
	MemberId   uint64 `json:"memberId"`
	ShopId     uint64 `json:"shopId"`
	MenuId     uint64 `json:"menuId"`
	OrderCount uint32 `json:"orderCount"`
	Category   string `json:"category"`
}
