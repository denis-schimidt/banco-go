package contas

import (
	"fmt"

	"github.com/denis-schimidt/banco/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorSolicitadoSaque float64) bool {
	podeSacar := valorSolicitadoSaque > 0 && c.saldo > valorSolicitadoSaque

	if podeSacar {
		c.saldo -= valorSolicitadoSaque
		fmt.Println("Saque de R$", valorSolicitadoSaque, "realizado com sucesso na conta do(a)", c.Titular)
		return true
	}

	if valorSolicitadoSaque < 0 {
		fmt.Println("Valor solicitado para saque R$", valorSolicitadoSaque, "é inválido")

	} else {
		fmt.Println("saldo insuficiente na conta do(a)", c.Titular, "para o saque solicitado de R$", valorSolicitadoSaque)
	}

	return false
}

func (c *ContaPoupanca) Depositar(valoresADepositar ...float64) (string, float64) {

	for _, valorDeposito := range valoresADepositar {

		if valorDeposito > 0 {
			c.saldo += valorDeposito
			fmt.Println("Depósito de R$", valorDeposito, "realizado com sucesso na conta do(a)", c.Titular)

		} else {
			fmt.Println("Valor de depósito inválido de R$", valorDeposito, "na conta do(a)", c.Titular)
		}
	}

	return fmt.Sprintf("%d depósitos processados", len(valoresADepositar)), c.saldo
}

func (c *ContaPoupanca) TransferirPara(ContaPoupancaDestino *ContaPoupanca, valorTransferencia float64) bool {
	saqueRealizado := c.Sacar(valorTransferencia)

	if saqueRealizado {
		ContaPoupancaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
