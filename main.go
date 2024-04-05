package main

import (
	"bufio"
	"fmt"
	"io"
	//"io/ioutil" - OBSOLETO
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const DELAY = 10

func main() {
	introduction()
	for {
		fmt.Println("")

		menu()

		comando := get()

		switch comando {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibindo logs...")
			readlogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando inesistente")
			os.Exit(-1)
		}
	}
}

func introduction() {
	nome := "Gabriel"
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Versão: ", versao)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func get() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func monitoring() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://www.alura.com.br", "https://www.github.com", "https://www.hackerrank.com"}

	sites := readsite()

	for i := 0; i < MONITORAMENTOS; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		fmt.Println("")
		writelogs(site, true)
	} else {
		fmt.Println("Site:", site, "apresenta problemas! StatusCode:", resp.StatusCode)
		fmt.Println("")
		writelogs(site, false)
	}
}

func readsite() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	read := bufio.NewReader(file)

	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func writelogs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - oline:" + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readlogs() {
	// file, err := ioutil.ReadFile("logs.txt") - OBSOLETO
	file, err := os.Open("logs.txt")

	read, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(read))

	file.Close()
}
