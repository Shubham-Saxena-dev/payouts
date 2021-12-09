package models

type Item struct {
	Name            string `json:"name" validate:"min=1,max=15,regexp=^[a-zA-Z]*$"`
	PriceAmount     int64  `json:"price_amount"`
	PriceCurrency   string `json:"price_currency"`
	SellerReference int32  `json:"seller_reference" validate:"min=1,max=999"`
}
