package service

import (
	"context"
)

type PrizeFetcher struct{}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}
