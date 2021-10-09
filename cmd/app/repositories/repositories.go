package repositories

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUserRepository)
	_ = container.Provide(NewCategoryRepository)
	_ = container.Provide(NewProductRepository)
	_ = container.Provide(NewQuantityRepository)
	_ = container.Provide(NewCartRepository)
	_ = container.Provide(NewOrderRepository)
	_ = container.Provide(NewOrderLineRepository)
	_ = container.Provide(NewAddressRepository)
	return nil
}
