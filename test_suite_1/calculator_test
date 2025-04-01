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
	if !t.Run("SUB Failure", func(t *testing.T) { assert.NotEqual(t, cal.Sub(10, 5), int32(5)) }) {
		t.Error("Test SUB Failure case failed")
	}
}
