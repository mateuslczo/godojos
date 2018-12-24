package main

import (
	"fmt"
	"strconv"
)

var currentBalance = 300
var notas = [...]int{100, 50, 50, 20, 20, 20, 20, 20, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

func main() {

	writeInTheScreen("Selecione o valor da retirada ou pressione \"S\" para sair: ")

	for currentBalance > 0 {

		serviceValue := requestMoney()

		if serviceValue == 0 {

			writeInTheScreen("Operação cancelada")
			break

		}

		if serviceValue > currentBalance {

			writeInTheScreen("Saldo insuficiente")

		} else {

			writeInTheScreen(releaseValue(serviceValue))

		}

	}
}

func releaseValue(serviceValue int) string {

	var holdValueToCompare string
	var totalSepareted int

	for i := 0; i < len(notas); i++ {

		if serviceValue == notas[i] {

			holdValueToCompare = strconv.Itoa(serviceValue)
			debitValue(notas[i])
			break

		}

	}

	for i := 0; i < len(notas); i++ {

		if serviceValue > notas[i] && totalSepareted <= serviceValue {

			if totalSepareted+notas[i] <= serviceValue {

				holdValueToCompare = holdValueToCompare + strconv.Itoa(notas[i]) + "; "
				totalSepareted = totalSepareted + notas[i]
				debitValue(notas[i])

			}

		}

		if serviceValue == notas[i] {
			break
		}

	}

	return holdValueToCompare

}

func debitValue(valueToDebit int) {

	currentBalance = currentBalance - valueToDebit

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
