package main

import (
	"fmt"
	"math"
)

func main() {
	var historial string
	var contador int
	for {
		var a, b float64
		var operacion string
		fmt.Println("==== CALCULADORA CIENTÍFICA v4.0 ====")
		fmt.Print("Ingresa el primer número: ")
		fmt.Scan(&a)
		fmt.Print("Ingresa el segundo número (si aplica): ")
		fmt.Scan(&b)
		fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
		fmt.Scan(&operacion)
		var resultadoTexto string
		switch operacion {
		case "+":
			resultado := a + b
			resultadoTexto = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, resultado)
		case "-":
			resultado := a - b
			resultadoTexto = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, resultado)
		case "*":
			resultado := a * b
			resultadoTexto = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, resultado)
		case "/":
			if b == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				continue
			}
			resultado := a / b
			resultadoTexto = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, resultado)
		case "^":
			resultado := math.Pow(a, b)
			resultadoTexto = fmt.Sprintf("%.2f ^ %.2f = %.2f", a, b, resultado)
		case "!":
			if a < 0 || a != float64(int(a)) {
				fmt.Println("Error: factorial solo para enteros positivos")
				continue
			}
			resultado := 1.0
			for i := 1; i <= int(a); i++ {
				resultado *= float64(i)
			}
			resultadoTexto = fmt.Sprintf("%.0f ! = %.0f", a, resultado)
		default:
			fmt.Println("Error: operación no reconocida")
			continue
		}
		//resultado
		fmt.Println("Resultado:", resultadoTexto)
		//historial
		contador++
		historial = historial + fmt.Sprintf("%d) %s\n", contador, resultadoTexto)
		//desea continuar?
		var opcion string
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scan(&opcion)
		if opcion == "n" {
			break
		}
	}
	//historial final
	fmt.Println("\n==== HISTORIAL DE OPERACIONES ====")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)
	fmt.Println("¡Hasta luego!")
}