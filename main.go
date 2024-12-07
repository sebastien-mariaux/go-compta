package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/expenses", getExpenses)

	router.Run("localhost:8080")
}

// Expense definition
type expense struct {
	ID     string  `json:"id"`
	Number   string  `json:"number"`
	Description  string  `json:"description"`
	GrossAmount  float64 `json:"gross_amount"`
	NetAmount  float64 `json:"net_amount"`
}

// Expense seeds
var expenses = []expense{
	{ID: "1", Number: "Ref001-43", Description: "Office Supplies", GrossAmount: 100.00, NetAmount: 80.00},
	{ID: "2", Number: "ABC-oct-2024", Description: "Energy", GrossAmount: 345.4, NetAmount: 312.2},
	{ID: "3", Number: "XXX0001", Description: "Car insurance", GrossAmount: 62.23, NetAmount: 56.98},
}

func getExpenses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, expenses)
}