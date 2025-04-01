package calculator

import "time"

type Calculator struct{}

func (*Calculator) Add(a, b int32) int32 {
	// doing some work
	time.Sleep(time.Millisecond * 50)
	return a + b
}

func (*Calculator) Sub(a, b int32) int32 {
	// doing some work
	time.Sleep(time.Millisecond * 50)
	return a - b
}

func (*Calculator) Mul(a, b int32) int32 {
	// doing some work
	time.Sleep(time.Millisecond * 50)
	return a * b
}

func (*Calculator) Div(a, b int32) float32 {
	// doing some work
	time.Sleep(time.Millisecond * 50)
	return float32(a / b)
}
