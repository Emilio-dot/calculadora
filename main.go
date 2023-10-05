package main

//Emilio Saldaña Salinas IX
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese la operación (por ejemplo, 2 + 3): ")
	scanner.Scan()
	operacion := scanner.Text()

	resultado, err := calcular(operacion)
	if err != nil {
		fmt.Println("Error al calcular:", err)
		return
	}

	fmt.Println("Resultado:", resultado)
}

func calcular(operacion string) (float64, error) {
	// Dividir la operación en operandos y operador
	parts := strings.Fields(operacion)
	if len(parts) != 3 {
		return 0, fmt.Errorf("formato de operación incorrecto")
	}

	// Convertir los operandos a números
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir el primer número: %v", err)
	}

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir el segundo número: %v", err)
	}

	// Realizar la operación según el operador
	var resultado float64
	switch parts[1] {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("división por cero no permitida")
		}
		resultado = num1 / num2
	default:
		return 0, fmt.Errorf("operador no válido")
	}

	return resultado, nil
}
