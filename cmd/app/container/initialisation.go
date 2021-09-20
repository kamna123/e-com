package container

import (
	"log"

	"e-commerce/cmd/app/repositories"

	"e-commerce/cmd/app/services"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	err := repositories.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject repositories", err)
	}
	err = services.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject services", err)
	}
	return container
}
