package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAmount = Amounts{
	ID: "1", NetAmount: 100.0, GrossAmount: 120.0,
}

func TestComputeVat(t *testing.T) {
	assert.Equal(t, 20.0, testAmount.ComputeVat())
}
