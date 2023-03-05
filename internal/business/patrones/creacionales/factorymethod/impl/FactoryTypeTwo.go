package impl

import (
	"example/patrones/internal/business/patrones/creacionales/factorymethod"
)

type factoryTypeTwo struct {
	factorymethod.Data
}

func NewFactoryTypeTwo() factorymethod.Factory {
	return &factoryTypeOne{
		factorymethod.Data{
			Info: "test 2",
		},
	}
}
