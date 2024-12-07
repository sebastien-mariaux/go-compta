package main

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/expenses", getExpenses)
	router.GET("/revenues", getRevenues)
	router.GET("/invoices", getInvoices)

	router.GET("/expenses/:id/vat", getExpenseVat)
	router.GET("/revenues/:id/vat", getRevenueVat)

	return router
}
