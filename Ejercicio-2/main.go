package main

import (
	"errors"
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
		return errors.New("el salario ingresado no alcanza el minimo imponible")
	}
	return nil
}
