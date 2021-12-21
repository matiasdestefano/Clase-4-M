package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	salario, err := calcularSalarioMensual(160, 1000.0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	aguinaldo, errAguinaldo := calcularMedioAguinaldo(130000, 6)
	if errAguinaldo != nil {
		fmt.Println(errAguinaldo.Error())
		os.Exit(1)
	}

	fmt.Printf("El salario es de: $%.2f\n", salario)
	fmt.Printf("El medio aguinaldo: $%.2f\n", aguinaldo)
}

func calcularSalarioMensual(horasTrabajadas int, valorHora float64) (float64, error) {
	var errorFinal error
	if horasTrabajadas < 80 {
		errorHoras := errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
		errorFinal = fmt.Errorf("%w", errorHoras)
		if horasTrabajadas < 0 {
			errorNegativo := fmt.Errorf("error: se ingreso un numero negativo para las horas %w", errorHoras)
			errorFinal = errorNegativo
		}
		return 0, errorFinal
	}

	salarioMensual := float64(horasTrabajadas) * valorHora
	if salarioMensual > 150000 {
		salarioMensual = salarioMensual * 0.9
	}

	return salarioMensual, nil
}

func calcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		return 0, fmt.Errorf("error: se ha ingresado un valor incorrecto para el salario o meses trabajados, [ %v ]", time.Now())
	}
	aguinaldo := mejorSalario / float64(12*mesesTrabajados)
	return aguinaldo, nil
}
