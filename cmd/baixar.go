package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Baixar(uri, filtroArquivo, arquivoSaida string) {

	// client := &http.Client{
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 		return http.ErrUseLastResponse
	// 	}}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("Cookie", "security=true")
	resp, err := client.Do(req)
	// resp, err := http.Get(uri) // http.Get(uri)

	// fmt.Println("StatusCode:", resp.StatusCode)
	// fmt.Println(resp.Request.URL)

	TratarErro("Não foi possivel baixar o arquivo:", err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	TratarErro("Não foi possivel ler o arquivo:", err)

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))

	TratarErro("Não foi possivel descompactar o arquivo:", err)

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		fmt.Println("Lendo arquivo: ", zipFile.Name)
		if zipFile.Name == filtroArquivo {
			unzippedFileBytes, err := LerZip(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			GravarArquivo(arquivoSaida, unzippedFileBytes)
		}
	}
}
