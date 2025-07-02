package main

import (
	"context"
	"log"
	"time"
)

type loggingService struct {
	next ICalculator
}

func NewLoggingService(next ICalculator) ICalculator {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) Calculate(ctx context.Context, a, b float32, operation string) (result float32, err error) {
	defer func(begin time.Time) {
		log.Printf("method=Calculate requestID=%v a=%f b=%f operation=%s result=%f err=%v took=%v",
			ctx.Value("requestID"), a, b, operation, result, err, time.Since(begin))
	}(time.Now())

	return s.next.Calculate(ctx, a, b, operation)
}
