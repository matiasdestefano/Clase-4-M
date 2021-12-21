package main

import (
	"fmt"
	"os"
)

type errorSalario struct {
	mensaje string
}

func main() {
	salary := 140000
	err := validarSalario(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func validarSalario(salary int) error {
	if salary < 150000 {
		return errorDeSalario()
	}
	return nil
}

func errorDeSalario() error {
	return &errorSalario{mensaje: "el salario ingresado no alcanza el mínimo imponible"}
}

func (e *errorSalario) Error() string {
	return fmt.Sprintf("Error: %s", e.mensaje)
}
