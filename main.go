package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hfantin/bolao-cli/cmd"
	"github.com/manifoldco/promptui"
)

func main() {
	exibirMenu()
	opcao := exibirPrompt("Opção")
	executarComando(opcao)
}

func exibirPrompt(label string) int {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validar,
	}
	opcao, err := prompt.Run()
	cmd.TratarErro("", err)
	opcaoInt, err := strconv.Atoi(opcao)
	cmd.TratarErro("", err)
	return opcaoInt
}

func exibirMenu() {
	fmt.Println("Seleciona a opção:")
	fmt.Println(" 1) Baixar Arquivos")
	fmt.Println(" 2) Gerar Arquivo csv")
	fmt.Println(" 3) Atualizar BD")
	fmt.Println(" 4) Gerar numeros aleatorios")
	fmt.Println(" 9) Sair \n")
}

func executarComando(opcao int) {
	switch opcao {
	case 1:
		cmd.Baixar("http://www1.caixa.gov.br/loterias/_arquivos/loterias/D_mgsasc.zip", "d_megasc.htm", "megasena.html")
	case 2:
		cmd.GerarCsv("arquivos/megasena.html", "resultados.csv")
	case 3:
		fmt.Println("atualizar bd - em construção...")
	case 4:
		dezenas := exibirPrompt("Informe a quantidade de dezenas")
		jogos := exibirPrompt("Informe a quantidade de jogos")
		cmd.GerarDezenas(dezenas, jogos)
	case 9:
		fmt.Println("Falou campeão!")
	default:
		fmt.Printf("Não existe a opção %v!\n.", opcao)
	}
}
func validar(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Número inválido")
	}
	return nil
}
