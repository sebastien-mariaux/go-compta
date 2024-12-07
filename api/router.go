package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/expenses", GetExpenses)
	router.GET("/revenues", GetRevenues)
	router.GET("/invoices", GetInvoices)

	router.GET("/expenses/:id/vat", GetExpenseVat)
	router.GET("/revenues/:id/vat", GetRevenueVat)

	return router
}
