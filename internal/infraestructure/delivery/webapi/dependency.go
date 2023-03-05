package webapi

import (
	"example/patrones/internal/infraestructure/delivery/webapi/handler"
)

type Dependencies struct {
	testHandler handler.TestHandler
}

func NewDependencies(testHandler handler.TestHandler) *Dependencies {
	return &Dependencies{
		testHandler: testHandler,
	}
}

func buildDependencies() *Dependencies {

	return &Dependencies{
		testHandler: handler.NewTestHandler(),
	}
}
