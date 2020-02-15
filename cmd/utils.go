package cmd

import (
	"fmt"
	"os"
)

func TratarErro(mensagem string, erro error) {
	if erro != nil {
		fmt.Println(mensagem, erro)
		os.Exit(1)
	}
}
