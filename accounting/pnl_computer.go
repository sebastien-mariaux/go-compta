package accounting

import (
	"fmt"
	"go-compta/models"
	"math"
)

func getTotalNetRevenue(data []models.Revenue ) float64 {
  var totalNetRevenue float64 = 0
  for _, revenue := range data {
    totalNetRevenue += revenue.NetAmount
  }

  return math.Round(totalNetRevenue*100)/100
}

func getTotalNetExpense(data []models.Expense) float64 {
  var totalNetExpense float64 = 0
  for _, expense := range data {
    totalNetExpense += expense.NetAmount
  }

  return math.Round(totalNetExpense*100)/100
}

func getCollectedVat(data []models.Revenue) float64 {
  var totalVat float64 = 0
  for _, revenue := range data {
    totalVat += revenue.ComputeVat()
  }

  return math.Round(totalVat*100)/100
}

func getDeductedVat(data []models.Expense) float64 {
  var totalVat float64 = 0
  for _, expense := range data {
    totalVat += expense.ComputeVat()
  }

  return math.Round(totalVat*100)/100
}

func ComputePNLBeforeTaxes(revenues []models.Revenue, expenses []models.Expense) float64 {
  var totalRevenues float64 = getTotalNetRevenue(revenues) - getTotalNetExpense(expenses)

  return math.Round(totalRevenues*100)/100
}

func ComputeVatToPay(revenues []models.Revenue, expenses []models.Expense) float64 {
  var vatToPay float64 = getCollectedVat(revenues) - getDeductedVat(expenses)

  return vatToPay
}

func ComputeTaxes(revenues []models.Revenue, expenses []models.Expense) float64 {
  var pnl = ComputePNLBeforeTaxes(revenues, expenses)
  fmt.Printf("%f\n", pnl)
  if pnl < 0 {
    return 0
  }
  var first_bracket float64 = math.Min(41500, pnl)
  var second_bracket float64 = math.Max(0, pnl - 41500)
  var taxes = 0.15 * first_bracket + 0.25 * second_bracket
  return math.Round(taxes*100)/100
}

func ComputeNetResult(revenues []models.Revenue, expenses []models.Expense) float64 {
  var pnl = ComputePNLBeforeTaxes(revenues, expenses)
  var taxes = ComputeTaxes(revenues, expenses)
  return math.Round((pnl - taxes)*100)/100
}