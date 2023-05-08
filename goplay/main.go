package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kwamekyeimonies/go-microservice/logger"
	"github.com/kwamekyeimonies/go-microservice/service"
)

func main() {
	svc := logger.NewLoggingService(&service.PrizeFetcher{})

	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}
