package main

import (
	"api/internal/application"
	"context"
	"fmt"
)

//	@title			Admin API
//	@version		1.0
//	@description	This is an API that handles the admin commands.
//	@termsOfService	https://www.linkedin.com/in/emilianokohmann/

//	@contact.name	Emiliano Kohmann
//	@contact.url	https://www.linkedin.com/in/emilianokohmann/
//	@contact.email	emikohmann@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:9999
//	@BasePath	/
//	@schemes	http
func main() {
	ctx := context.Background()

	app, err := application.NewApplication(ctx)
	if err != nil {
		panic(fmt.Errorf("error creating application: %w", err))
	}

	if err := app.Run(ctx); err != nil {
		panic(fmt.Errorf("error running application: %w", err))
	}
}
