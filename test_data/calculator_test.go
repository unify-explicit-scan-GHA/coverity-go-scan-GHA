package suite_1

import (
	"os"
	"testing"

	"github.com/calculi-corp/template-go-testing/calculator"
	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) != 0 {
		t.Skip("skiping this test case")
	}
	t.Log("GOPATH: ", os.Getenv("GOPATH"))
}

func TestAdd(t *testing.T) {

	// time.Sleep(time.Millisecond * 50)
	cal := calculator.Calculator{}

	if !t.Run("ADD Success", func(t *testing.T) { assert.Equal(t, cal.Add(10, 10), int32(20)) }) {
		t.Error("Test ADD Success case failed")
	}
	if !t.Run("ADD Failure", func(t *testing.T) { assert.NotEqual(t, cal.Add(10, 10), int32(10)) }) {
		t.Error("Test ADD Failure case failed")
	}
}
func TestSub(t *testing.T) {

	// time.Sleep(time.Millisecond * 50)
	cal := calculator.Calculator{}

	if !t.Run("SUB Success", func(t *testing.T) { assert.Equal(t, cal.Sub(10, 5), int32(5)) }) {
		t.Error("Test SUB Success case failed")
	}
}

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
}
