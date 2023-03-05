package services

import (
	"example/patrones/internal/business/domain"
	"example/patrones/internal/business/patrones/creacionales/factorymethod"
	"example/patrones/internal/business/patrones/creacionales/factorymethod/impl"
)

type FactoryService struct {
}

func (f *FactoryService) Test(typefactory string) factorymethod.Factory {
	switch typefactory {
	case string(domain.TypeOne):
		return impl.NewFactoryTypeOne()

	case string(domain.TypeTwo):
		return impl.NewFactoryTypeTwo()

	default:
		return nil
	}
}
