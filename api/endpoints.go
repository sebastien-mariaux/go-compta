package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-compta/accounting"
	"go-compta/models"
)

func GetExpenses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.ExpensesData)
}

func GetRevenues(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.RevenuesData)
}

func GetInvoices(c *gin.Context) {
	var invoices []interface{}
	for _, expense := range models.ExpensesData {
		invoices = append(invoices, expense)
	}
	for _, revenue := range models.RevenuesData {
		invoices = append(invoices, revenue)
	}
	c.IndentedJSON(http.StatusOK, invoices)
}

func GetExpenseVat(context *gin.Context) {
	id := context.Param("id")
	for _, expense := range models.ExpensesData {
		if expense.ID == id {
			context.IndentedJSON(http.StatusOK, models.Amounts{
				ID:          expense.ID,
				NetAmount:   expense.NetAmount,
				GrossAmount: expense.GrossAmount,
			})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense not found"})
}

func GetRevenueVat(context *gin.Context) {
	id := context.Param("id")
	for _, revenue := range models.RevenuesData {
		if revenue.ID == id {
			context.IndentedJSON(http.StatusOK, models.Amounts{
				ID:          revenue.ID,
				NetAmount:   revenue.NetAmount,
				GrossAmount: revenue.GrossAmount,
			})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "revenue not found"})
}

func GetNetPNL(context *gin.Context) {
	var netPNL = accounting.ComputeNetResult(models.RevenuesData, models.ExpensesData)
	context.IndentedJSON(http.StatusOK, gin.H{"netPNL": netPNL})
}

func GetPNL(context *gin.Context) {
	var pnl = accounting.ComputePNLBeforeTaxes(models.RevenuesData, models.ExpensesData)
	context.IndentedJSON(http.StatusOK, gin.H{"pnl": pnl})
}