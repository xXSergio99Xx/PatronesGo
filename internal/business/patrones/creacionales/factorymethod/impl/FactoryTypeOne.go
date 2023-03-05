package impl

import (
	"example/patrones/internal/business/patrones/creacionales/factorymethod"
)

type factoryTypeOne struct {
	factorymethod.Data
}

func NewFactoryTypeOne() factorymethod.Factory {
	return &factoryTypeOne{
		factorymethod.Data{
			Info: "test 1",
		},
	}
}
