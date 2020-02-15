package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"time"
)

const tdRegex = "<td rowspan=\"1\">(.*)</td>"
const trOpenRegex = "<tr(.*)>"
const trCloseRegex = "</tr>"

func GerarCsv(arquivoEntrada, arquivoSaida string) {
	file, err := os.Open(arquivoEntrada)
	TratarErro("não foi possível abrir o arquivo", err)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	matchTrOpen := false
	matchTrClose := false
	matchTd := regexp.MustCompile(tdRegex)
	texto := ""
	colunas := 0
	for scanner.Scan() {
		if !matchTrOpen {
			matchTrOpen, _ = regexp.MatchString(trOpenRegex, scanner.Text())
		}
		valores := matchTd.FindStringSubmatch(scanner.Text())
		if matchTrOpen && len(valores) > 0 && colunas < 8 {
			// verifica se o valor é uma da
			data, err := time.Parse("02/01/2006", valores[1])
			if err != nil {
				texto += valores[1] + ";"
			} else {
				texto += data.Format("2006-01-02") + ";"
			}
			colunas++
		}

		matchTrClose, _ = regexp.MatchString(trCloseRegex, scanner.Text())
		if matchTrClose {
			if colunas > 0 {
				texto += "\n"
			}
			colunas = 0
			matchTrOpen = false
			matchTrClose = false
		}

	}
	GravarArquivo(arquivoSaida, []byte(texto))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
