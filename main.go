package main

import (
	"go-compta/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupRouter().Run("localhost:8080")
}

func getExpenses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, expenses)
}

func getRevenues(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, revenues)
}

func getInvoices(c *gin.Context) {
	var invoices []interface{}
	for _, expense := range expenses {
		invoices = append(invoices, expense)
	}
	for _, revenue := range revenues {
		invoices = append(invoices, revenue)
	}
	c.IndentedJSON(http.StatusOK, invoices)
}

func getExpenseVat(context *gin.Context) {
	id := context.Param("id")
	for _, expense := range expenses {
		if expense.ID == id {
			context.IndentedJSON(http.StatusOK, models.Amounts{
				ID: expense.ID,
				NetAmount: expense.NetAmount,
				GrossAmount: expense.GrossAmount,
			})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense not found"})
}

func getRevenueVat(context *gin.Context) {
	id := context.Param("id")
	for _, revenue := range revenues {
		if revenue.ID == id {
			context.IndentedJSON(http.StatusOK, models.Amounts{
				ID: revenue.ID,
				NetAmount: revenue.NetAmount,
				GrossAmount: revenue.GrossAmount,
			})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "revenue not found"})
}

