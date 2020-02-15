package cmd

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

func GravarArquivo(name string, data []byte) {
	output, err := os.Create("arquivos/" + name)
	TratarErro("Nao foi possivel criar o arquivo de saida: ", err)
	defer output.Close()

	if _, err := output.Write(data); err != nil {
		fmt.Println("Nao foi possivel gravar o arquivo de saida: ", err)
		os.Exit(1)
	}

	fmt.Println("Arquivo gravado com sucesso: ", name)
}

func LerZip(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
