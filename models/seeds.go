package models

// Expense seeds
var ExpensesData = []Expense{
	{
		Invoice: Invoice{
			ID:          "1",
			Number:      "Ref001-43",
			Description: "Office Supplies",
			GrossAmount: 100.00,
			NetAmount:   80.00,
		},
		Supplier: "Super-office",
		Category: "Expense",
	},
	{
		Invoice: Invoice{
			ID:          "2",
			Number:      "ABC-oct-2024",
			Description: "Energy",
			GrossAmount: 345.4,
			NetAmount:   312.2,
		},
		Supplier: "Engie",
		Category: "Expense",
	},
	{
		Invoice: Invoice{
			ID:          "3",
			Number:      "XXX0001",
			Description: "Car insurance",
			GrossAmount: 62.23,
			NetAmount:   56.98,
		},
		Supplier: "AXA",
		Category: "Expense",
	},
}

// Revenue seeds
var RevenuesData = []Revenue{
	{
		Invoice: Invoice{
			ID:          "4",
			Number:      "CLI-0001",
			Description: "Consulting",
			GrossAmount: 1000.00,
			NetAmount:   800.00,
		},
		Customer: "Help-me-corp",
		Category: "Revenue",
	},
	{

		Invoice: Invoice{
			ID:          "5",
			Number:      "DEV-0001",
			Description: "Development",
			GrossAmount: 3450.4,
			NetAmount:   3122.2,
		},
		Customer: "wedev.org",
		Category: "Revenue",
	},
	{

		Invoice: Invoice{
			ID:          "6",
			Number:      "XXX0001",
			Description: "Training",
			GrossAmount: 622.23,
			NetAmount:   569.98,
		},
		Customer: "train-me-quick",
		Category: "Revenue",
	},
}
