package request

type Message struct {
	MemberId uint64 `json:"memberId"`
	ShopId   uint64 `json:"shopId"`
	MenuId   uint64 `json:"menuId"`
}
