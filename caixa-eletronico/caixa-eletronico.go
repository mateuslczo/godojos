package main

import (
	"fmt"
	"strconv"
)

var currentBalance = 300
var notas = [...]int{100, 50, 50, 20, 20, 20, 20, 20, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

func main() {

	writeInTheScreen("Selecione o valor da retirada: ")

	/*
		for i := 0; i < len(notas); i++ {

			writeInTheScreen(strconv.Itoa(notas[i]) + "; ")

		}
	*/

	var serviceValue = requestMoney()
	if serviceValue > currentBalance {
		writeInTheScreen("Saldo insuficiente")
	} else {
		writeInTheScreen(releaseValue(serviceValue))
	}

}

func releaseValue(serviceValue int) string {

	var holdValueToCompare string
	var totalSepareted int

	for i := 0; i < len(notas); i++ {

		if serviceValue == notas[i] {
			holdValueToCompare = strconv.Itoa(serviceValue)
			break
		}

	}

	for i := 0; i < len(notas); i++ {

		if serviceValue > notas[i] && totalSepareted <= serviceValue {

			if totalSepareted+notas[i] <= serviceValue {
				holdValueToCompare = holdValueToCompare + strconv.Itoa(notas[i]) + "; "
				totalSepareted = totalSepareted + notas[i]
			}

		}

		if serviceValue == notas[i] {
			break
		}

	}

	return holdValueToCompare

}

/*
método para escrever na tela
*/
func writeInTheScreen(word string) {

	fmt.Printf(word + "\n")

}

/*
método retornar saldo
*/
func queryBalance(withdrawalAmount int) bool {

	if withdrawalAmount <= currentBalance {
		return true
	}
	return false

}

/*
 Método para ler valor do teclado
*/
func requestMoney() int {

	var withdrawalAmount int

	fmt.Scan(&withdrawalAmount)

	return withdrawalAmount

}
