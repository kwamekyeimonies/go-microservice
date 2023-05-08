package logger

import (
	"context"
	"time"

	"github.com/kwamekyeimonies/go-microservice/service"
	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next service.PriceFetcher
}

func NewLoggingService(next service.PriceFetcher) service.PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(
			logrus.Fields{
				"took":  time.Since(begin),
				"err":   err.Error(),
				"price": price,
			},
		).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
