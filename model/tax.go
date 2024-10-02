package model

type CalculateTaxRequest struct {
	LineItems       []ProductLineItem `json:"line_items"`
	CustomerDetails CustomerDetails   `json:"customer_details"`
	ShippingFee     float64           `json:"shipping_fee"`
}

type CalculateTaxResponse struct {
	ID          string  `json:"id"`
	AmountTotal float64 `json:"amount_total"`
	TaxAmount   float64 `json:"tax_amount"`
}

type ProductLineItem struct {
	Amount      float64 `json:"amount"`
	Quantity    int64   `json:"quantity"`
	Reference   string  `json:"reference"`
	TaxCode     string  `json:"tax_code"`
	TaxBehavior string  `json:"tax_behavior"`
}

type CustomerDetails struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	AddressType string `json:"address_type"`
	City        string `json:"city"`
	Line1       string `json:"line1"`
}
