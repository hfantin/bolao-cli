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
const MAX_COL = 10
const LINHA_REGEX = "([0-9]+);[.]*"
const DATA_SQL_PATTERN = "2006-01-02"
const DIA_MES_ANO_PATTERN = "02/01/2006"

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
		if matchTrOpen && len(valores) > 0 && colunas < MAX_COL {
			colunas++
			if colunas == 9 { // coluna de arrecadacao
				continue
			}

			// verifica se o valor é uma da
			data, err := time.Parse(DIA_MES_ANO_PATTERN, valores[2])
			if err != nil {
				linha += valores[2] + ";"
			} else {
				linha += data.Format(DATA_SQL_PATTERN) + ";"
			}

		}

		matchTrClose, _ = regexp.MatchString(TR_CLOSE_REGEX, scanner.Text())
		if matchTrClose {
			match, _ := regexp.MatchString(LINHA_REGEX, linha)
			if colunas > 0 && match {
				texto += linha[:len(linha)-1] + "\n"
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
