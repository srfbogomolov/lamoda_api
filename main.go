package main

import (
	"context"
	"log"

	"github.com/srfbogomolov/warehouse_api/config"
)

func main() {
	_, err := config.NewConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
