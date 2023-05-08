package service

import (
	"context"
	"fmt"
)

func (s *PrizeFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not Supported", ticker)
	}

	return price, nil
}
