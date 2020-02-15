package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func GerarDezenas(qtd, jogos int) {
	max := 60

	if qtd < 6 {
		qtd = 6
	}

	if jogos < 1 {
		jogos = 1
	}

	fmt.Printf("Gerando bolão da megasena com %d numeros e %d jogo(s):\n\n", qtd, jogos)
	for i := 0; i < jogos; i++ {
		numeros := gerarNumeros(max, qtd)
		fmt.Printf("#%d: %v\n", i+1, numeros)
	}
}

func getArgAsInt(valor, nome string) int {
	convertido, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Printf("Campo %s invalido: \"%s\"\n", nome, valor)
		os.Exit(0)
	}
	return convertido
}

func gerarNumeros(max, qtd int) []int {
	rand.Seed(time.Now().UnixNano())
	mapa := make(map[int]struct{})
	for i := 0; i < qtd && len(mapa) < max; i++ {
		var numero int
		exists := true
		for exists {
			numero = rand.Intn(max) + 1
			_, exists = mapa[numero]
		}
		mapa[numero] = struct{}{}
	}
	slice := extrairNumeros(mapa)
	sort.Ints(slice)
	return slice
}

func extrairNumeros(mapa map[int]struct{}) []int {
	slice := []int{}
	for k := range mapa {
		slice = append(slice, k)
	}
	return slice
}
