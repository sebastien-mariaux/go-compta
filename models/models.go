package models

import "encoding/json"

type Invoice struct {
	ID          string  `json:"id"`
	Number      string  `json:"number"`
	Description string  `json:"description"`
	GrossAmount float64 `json:"grossAmount"`
	NetAmount   float64 `json:"netAmount"`
}

func (i Invoice) ComputeVat() float64 {
	return i.GrossAmount - i.NetAmount
}

// Expense definition
type Expense struct {
	Invoice
	Supplier string `json:"supplier"`
	Category string `json:"category"`
}

func (a Amounts) ComputeVat() float64 {
	return a.GrossAmount - a.NetAmount
}

// Revenue definition
type Revenue struct {
	Invoice
	Customer string `json:"customer"`
	Category string `json:"category"`
}

// Amounts
type Amounts struct {
	ID          string  `json:"id"`
	NetAmount   float64 `json:"netAmount"`
	GrossAmount float64 `json:"grossAmount"`
}

// Custom JSON marshaller for amounts
func (a Amounts) MarshalJSON() ([]byte, error) {
	type Alias Amounts
	return json.Marshal(&struct {
		VAT float64 `json:"vat"`
		Alias
	}{
		VAT:   a.ComputeVat(),
		Alias: (Alias)(a),
	})
}