### ferramenta de linha de comando para o Bolao

### criar modulo

go mod init github.com/hfantin/bolao-cli

### Bibliotecas utilizadas

- cobra:
  > go get -u github.com/spf13/cobra/cobra

### avaliar bibliotecas

- viper
- cli

### links

[5 keys to go cli](https://blog.alexellis.io/5-keys-to-a-killer-go-cli/)

### criar prompt

[promptui](https://github.com/manifoldco/promptui)
[code overview](https://golang.org/doc/code.html#Overview)

### problemas com o caractere abaixo
"�" \uFFFD - replacement char
tentar strings.ToValidUTF8("a\xc5z", "")
