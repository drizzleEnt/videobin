package main

import (
	"context"
	"log"
	"videobin/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed run app: %s", err.Error())
	}
}
