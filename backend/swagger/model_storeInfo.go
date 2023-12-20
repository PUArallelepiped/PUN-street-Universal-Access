package swagger

type StoreInfo struct {
	StoreId int64 `json:"store_id"`

	Description string `json:"description"`

	Name string `json:"name"`

	Address string `json:"address"`

	Rate float32 `json:"rate"`

	RateCount int64 `json:"rate_count"`

	Picture string `json:"picture"`

	Status int64 `json:"status"`

	ShippingFee int64 `json:"shipping_fee"`
}
