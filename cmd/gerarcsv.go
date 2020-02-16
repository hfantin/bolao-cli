package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"time"
)

const TD_REGEX = "<td(.*)>(.*)</td>"
const TR_OPEN_REGEX = "<tr(.*)>"
const TR_CLOSE_REGEX = "</tr>"

func GerarCsv(arquivoEntrada, arquivoSaida string) {
	file, err := os.Open(arquivoEntrada)
	TratarErro("não foi possível abrir o arquivo", err)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	matchTrOpen := false
	matchTrClose := false
	matchTd := regexp.MustCompile(TD_REGEX)
	linha := ""
	texto := ""
	colunas := 0
	for scanner.Scan() {
		if !matchTrOpen {
			matchTrOpen, _ = regexp.MatchString(TR_OPEN_REGEX, scanner.Text())
		}
		valores := matchTd.FindStringSubmatch(scanner.Text())
		if matchTrOpen && len(valores) > 0 && colunas < 8 {
			// verifica se o valor é uma da
			data, err := time.Parse("02/01/2006", valores[2])
			if err != nil {
				linha += valores[2] + ";"
			} else {
				linha += data.Format("2006-01-02") + ";"
			}
			colunas++
		}

		matchTrClose, _ = regexp.MatchString(TR_CLOSE_REGEX, scanner.Text())
		if matchTrClose {
			match, _ := regexp.MatchString("([0-9]+);[.]*", linha)
			if colunas > 0 && match {
				texto += linha + "\n"
			}
			linha = ""
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
