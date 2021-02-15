package main

import (
	"fmt"

	"github.com/denis-schimidt/banco/clientes"
	"github.com/denis-schimidt/banco/contas"
)

func pagarBoleto(conta contaDebitavel, valorBoleto float64) bool {
	return conta.Sacar(valorBoleto)
}

type contaDebitavel interface {
	Sacar(valorSolicitadoSaque float64) bool
}

func main() {
	fulano := clientes.Titular{Nome: "Fulano", CPF: "264.084.968-87", Profissao: "Desenvolvedor"}
	contaFulano := contas.ContaPoupanca{Titular: fulano, NumeroAgencia: 123, NumeroConta: 1233456}
	contaFulano.Depositar(1000)

	ciclano := clientes.Titular{Nome: "Ciclano", CPF: "263.094.162-52", Profissao: "Desenvolvedor"}
	contaCiclano := contas.ContaPoupanca{Titular: ciclano, NumeroAgencia: 123, NumeroConta: 9999999}

	fmt.Println("Tranferência realizada:", contaFulano.TransferirPara(&contaCiclano, 100))

	fmt.Println("Saldo Fulano R$", contaFulano.ObterSaldo())
	fmt.Println("Saldo Ciclano R$", contaCiclano.ObterSaldo())

	pagarBoleto(&contaFulano, 150.0)
	fmt.Println("Saldo Fulano depois do boleto R$", contaFulano.ObterSaldo())

	beltrano := clientes.Titular{Nome: "Beltrano", CPF: "856.142.986-22", Profissao: "Gerente"}
	contaBeltrano := contas.ContaCorrente{Titular: beltrano, NumeroAgencia: 123, NumeroConta: 741256}
	contaBeltrano.Depositar(20000)
	fmt.Println("Saldo Beltrano R$", contaFulano.ObterSaldo())

	pagarBoleto(&contaBeltrano, 966.36)
	fmt.Println("Saldo Beltrano depois do boleto R$", contaBeltrano.ObterSaldo())
}
