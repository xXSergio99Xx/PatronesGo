package handler

import (
	"example/patrones/internal/business/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler interface {
	GetFactory(c *gin.Context)
}

type testHandler struct {
	// Poner casos de uso
	factoryService services.FactoryService
}

func NewTestHandler() *testHandler {
	return &testHandler{}
}

func (t testHandler) GetFactory(c *gin.Context) {
	types := c.Query("type")
	v := t.factoryService.Test(types)
	v.DoFactory()
	c.IndentedJSON(http.StatusOK, "hola mundo")
}
