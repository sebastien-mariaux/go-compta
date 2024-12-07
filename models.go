package main

import "encoding/json"

type invoice struct {
	ID          string  `json:"id"`
	Number      string  `json:"number"`
	Description string  `json:"description"`
	GrossAmount float64 `json:"grossAmount"`
	NetAmount   float64 `json:"netAmount"`
}

// Expense definition
type expense struct {
	invoice
	Supplier string `json:"supplier"`
	Category string `json:"category"`
}


// Revenue definition
type revenue struct {
	invoice
	Customer string `json:"customer"`
	Category string `json:"category"`
}

// Amounts
type amounts struct {
	ID 				string  `json:"id"`
	NetAmount   float64 `json:"netAmount"`
	GrossAmount float64 `json:"grossAmount"`
}

// Custom JSON marshaller for amounts
func (a amounts) MarshalJSON() ([]byte, error) {
	type Alias amounts
	return json.Marshal(&struct {
			VAT float64 `json:"vat"`
			Alias
	}{
			VAT:   a.GrossAmount - a.NetAmount,
			Alias: (Alias)(a),
	})
}