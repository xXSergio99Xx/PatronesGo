package factorymethod

import "fmt"

type Data struct {
	Info string
}

func (d *Data) DoFactory() {
	fmt.Print("Se imprime ", d.Info)
}
