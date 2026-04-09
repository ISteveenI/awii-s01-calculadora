package main

import "fmt"

func main() {

	var a, b float64
	var operacion string
	fmt.Println("==== CALCULADORA CIENTÍFICA v2.0 ====")
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&b)
	fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
	fmt.Scan(&operacion)
	switch operacion {
	case "+":
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: no se puede dividir entre cero")
			return
		}
		fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", a, b, a/b)
	case "^":
		if b < 0 {
			fmt.Println("Error: el exponente debe ser positivo")
			return
		}
		base := a
		resultado := 1.0
		for i := 0; i < int(b); i++ {
			resultado *= base
		}
		fmt.Printf("Resultado: %.2f ^ %.0f = %.2f\n", a, b, resultado)
	case "!":
		if a < 0 {
			fmt.Println("Error: no existe factorial de números negativos")
			return
		}
		resultado := 1.0
		for i := 1; i <= int(a); i++ {
			resultado *= float64(i)
		}
		fmt.Printf("Resultado: %.0f! = %.0f\n", a, resultado)
	default:
		fmt.Println("Error: operación no reconocida")
	}
}