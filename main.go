package main

import "fmt"

func main() {
	// Msg
	fmt.Println("[Dependendo da quantidade de paginas a serem scrapeadas isso pode demorar bastante]")

	/*
		Essa função precisa de dois parametros
		(total de paginas a serem scrapeadas ,e o nome do arquivo de saida(sem extensão))
		maximo que pode ser scrapeada é 347 ou o maximo de paginação que esta no site
	*/
	StartScraper(1, "json_db")
}
