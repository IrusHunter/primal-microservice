package main

import (
	"context"
	"fmt"
	"time"
)

type ICalculator interface {
	Calculate(ctx context.Context, a, b float32, operation string) (float32, error)
}

type DummyCalculator struct{}

func (dc DummyCalculator) Calculate(ctx context.Context, a, b float32, operation string) (float32, error) {
	time.Sleep(100 * time.Millisecond)

	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
