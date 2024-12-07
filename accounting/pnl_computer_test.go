package accounting

import (
	"go-compta/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalNetRevenue(t *testing.T) {
  assert.Equal(t, 4492.18, getTotalNetRevenue(models.RevenuesData))
}


func TestGetTotalNetExpense(t *testing.T) {
  assert.Equal(t, 449.18, getTotalNetExpense(models.ExpensesData))
}

func TestGetCollectedVat(t *testing.T) {
  assert.Equal(t, 580.45, getCollectedVat(models.RevenuesData))
}

func TestGetDeductedVat(t *testing.T) {
  assert.Equal(t, 58.45, getDeductedVat(models.ExpensesData))
}

func TestComputePNLBeforeTaxes(t *testing.T) {
  assert.Equal(t, 4043.0, ComputePNLBeforeTaxes(models.RevenuesData, models.ExpensesData))
}

func TestComputeVatToPay(t *testing.T) {
  assert.Equal(t, 522.0, ComputeVatToPay(models.RevenuesData, models.ExpensesData))
}

func TestComputeTaxesWhenLowRevenue(t *testing.T) {
  var revenues = []models.Revenue{
    {Invoice:models.Invoice{ID: "1", NetAmount: 21000.0, GrossAmount: 23000.0}},
  }

  var expenses = []models.Expense{
    {Invoice:models.Invoice{ID: "2", NetAmount: 1000.0, GrossAmount: 1200.0}},
  }
  assert.Equal(t, 0.15 * 20000, ComputeTaxes(revenues, expenses))
}

func TestComputeTaxesWhenNegativeRevenue(t *testing.T) {
  var revenues = []models.Revenue{
    {Invoice:models.Invoice{ID: "1", NetAmount: 21000.0, GrossAmount: 23000.0}},
  }

  var expenses = []models.Expense{
    {Invoice:models.Invoice{ID: "2", NetAmount: 100000.0, GrossAmount: 120000.0}},
  }
  assert.Equal(t, 0.0, ComputeTaxes(revenues, expenses))
}

func TestComputeTaxesWhenHighRevenue(t *testing.T) {
  var revenues = []models.Revenue{
    {Invoice:models.Invoice{ID: "1", NetAmount: 210000.0, GrossAmount: 230000.0}},
  }

  var expenses = []models.Expense{}
  // 0.15 * 41500
  // + 0.25 * (210000 - 41500)
  assert.Equal(t, 48350.0, ComputeTaxes(revenues, expenses))
}

func TestComputeNetResult(t *testing.T) {
  var revenues = []models.Revenue{
    {Invoice:models.Invoice{ID: "1", NetAmount: 210000.0, GrossAmount: 230000.0}},
  }

  var expenses = []models.Expense{}
  // 210000 - 0 - 48350
  assert.Equal(t, 161650.0, ComputeNetResult(revenues, expenses))
}