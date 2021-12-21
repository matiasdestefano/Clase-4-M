package main

import (
	"fmt"
	"os"
)

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
		return errorDeSalario(salary)
	}
	return nil
}

func errorDeSalario(salario int) error {
	return fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresao es de: %d", salario)
}
