package contas

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorSolicitadoSaque float64) bool {
	podeSacar := valorSolicitadoSaque > 0 && c.saldo > valorSolicitadoSaque

	if podeSacar {
		c.saldo -= valorSolicitadoSaque
		fmt.Println("Saque de R$", valorSolicitadoSaque, "realizado com sucesso na conta do(a)", c.Titular)
		return true
	}

	fmt.Println("Saldo insuficiente na conta do(a)", c.Titular, "para o saque solicitado de R$", valorSolicitadoSaque)
	return false
}

func (c *ContaCorrente) depositar(valoresADepositar ...float64) (string, float64) {

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

func (c *ContaCorrente) transferirPara(contaCorrenteDestino *ContaCorrente, valorTransferencia float64) bool {
	saqueRealizado := c.sacar(valorTransferencia)

	if saqueRealizado {
		contaCorrenteDestino.depositar(valorTransferencia)
		return true
	}

	fmt.Println("Saque de R$ ", valorTransferencia, "não foi realizado")
	return false
}
