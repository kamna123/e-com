package api

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserAPI)
	_ = container.Provide(NewCategoryAPI)
	_ = container.Provide(NewProductAPI)
	_ = container.Provide(NewQuantityAPI)
	_ = container.Provide(NewCartAPI)
	_ = container.Provide(NewOrderAPI)
	return nil
}
