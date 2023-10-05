package operadores

import (
	"strconv"
	"strings"
)

func Sumar(operacion string) int {
	valores := strings.Split(operacion, "+")
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, _ := strconv.Atoi(valores[i])
		resultado += num
	}

	return resultado
}
