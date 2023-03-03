package main

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/application"
)

// @title Items API
// @version 1.0
// @description This is an API that handles the items information.
// @termsOfService https://www.linkedin.com/in/emilianokohmann/

// @contact.name Emiliano Kohmann
// @contact.url https://www.linkedin.com/in/emilianokohmann/
// @contact.email emikohmann@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	app, err := application.NewApplication()
	if err != nil {
		panic(fmt.Errorf("error creating application: %w", err))
	}

	if err := app.Run(); err != nil {
		panic(fmt.Errorf("error running application: %w", err))
	}
}
