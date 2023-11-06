package main

import (
	"contas/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaExemplo := contas.ContaPoupanca{}
	contaDoLuiz := contas.ContaCorrente{}

	contaDoLuiz.Depositar(592)
	contaExemplo.Depositar(500)

	PagarBoleto(&contaDoLuiz, 60)
	PagarBoleto(&contaExemplo, 20)
	fmt.Println(contaDoLuiz.ObterSaldo())
	fmt.Println(contaExemplo.ObterSaldo())
}
