package models

type Payout struct {
	SellerReference int32  `json:"seller_reference"`
	Amount          int64  `json:"amount"`
	Currency        string `json:"currency"`
}

type PayoutKey struct {
	SellerReference int32
	Currency        string
}

type ApiResponse struct {
	NoOfTransactions int      `json:"no_of_transactions"`
	Payout           []Payout `json:"payouts"`
}
