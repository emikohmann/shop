package main

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/application"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		panic(fmt.Errorf("error creating application: %w", err))
	}

	if err := app.Run(); err != nil {
		panic(fmt.Errorf("error running application: %w", err))
	}
}
