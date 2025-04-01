package suite_2

import (
	"testing"

	"github.com/calculi-corp/template-go-testing/calculator"
	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {

	// time.Sleep(time.Millisecond * 50)
	cal := calculator.Calculator{}

	if !t.Run("MUL Success", func(t *testing.T) { assert.Equal(t, cal.Mul(10, 10), int32(100)) }) {
		t.Error("Test MUL Success case failed")
	}
	if !t.Run("MUL Failure", func(t *testing.T) { assert.NotEqual(t, cal.Mul(10, 10), int32(50)) }) {
		t.Error("Test MUL Failure case failed")
	}
}

func TestDiv(t *testing.T) {

	// time.Sleep(time.Millisecond * 50)
	cal := calculator.Calculator{}

	if !t.Run("DIV Success", func(t *testing.T) { assert.Equal(t, cal.Div(10, 2), float32(5)) }) {
		t.Error("Test DIV Success case failed")
	}
	if !t.Run("DIV Failure", func(t *testing.T) { assert.NotEqual(t, cal.Div(10, 2), float32(5)) }) {
		t.Error("Test DIV Failure case failed")
	}
}
