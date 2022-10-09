package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
)

// Structs
type SMangaInfoPageSource struct {
	Title       string
	Gender      string
	Author      string
	Artist      string
	Status      string
	Description string
	Captules    TCaptuleWrapper
}

type SMangaCaptulesInfo struct {
	CaptuleNumber string
	DatePublish   string
	CaptulesPage  []string
}

// Types
type Routes []string
type InfoManga []SMangaInfoPageSource
type TCaptuleWrapper []SMangaCaptulesInfo

// Globals
var RouteToManga Routes
var BasicInfoManga InfoManga

func get_routes(scraperPerPage *int) {
	// Quantidade de paginas a serem scrapeadas - max 347
	total_per_page := *scraperPerPage

	for i := 1; i <= total_per_page; i++ {
		// Acessa a pagina
		load_url := "https://goldenmangas.top/mangas&pagina=" + fmt.Sprint(i)

		html_page, err := htmlquery.LoadURL(load_url)
		if err != nil {
			panic(err)
		}

		//fmt.Println("-------------------------------------------------------\n\n\n") // debug

		// Pega a rota para o manga
		routes, err := htmlquery.QueryAll(html_page, "//article/div[2]/section/div/a/@href")
		if err != nil {
			panic(err)
		}

		for _, r := range routes {
			route_to_str := "https://goldenmangas.top" + htmlquery.InnerText(r)
			RouteToManga = append(RouteToManga, route_to_str)
		}

	}

}

// Pega informações sobre o manga como titulo e etc...
func RenderScraper() {

	for _, url := range RouteToManga {
		/// msg
		fmt.Println("Scraper em andamento....")

		mangaPageSource, err := htmlquery.LoadURL(url)

		if err != nil {
			panic(err)
		}

		raw_title := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/h2[1]")
		raw_gender := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/h5")
		raw_author := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/h5[2]")
		raw_artist := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/h5[3]")
		raw_status := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/h5[4]")
		raw_description := htmlquery.FindOne(mangaPageSource, "//article/div[2]/div[2]/div[1]/div[1]/div[2]/div[2]")

		title := htmlquery.InnerText(raw_title)
		gender := htmlquery.InnerText(raw_gender)
		author := htmlquery.InnerText(raw_author)
		artist := htmlquery.InnerText(raw_artist)
		status := htmlquery.InnerText(raw_status)
		description := htmlquery.InnerText(raw_description)

		// Acessando a elementos dos capitulos
		captule_wrapper, err := htmlquery.QueryAll(mangaPageSource, "//article/div[2]/div[2]/div[1]/ul/li")
		if err != nil {
			panic(err)
		}

		// Pega imagems
		var CaptuleWrapper TCaptuleWrapper
		var datePublish string
		for _, captule := range captule_wrapper {

			// Link do capitulo

			linkCaptule := "https://goldenmangas.top" + htmlquery.InnerText(htmlquery.FindOne(captule, "//a/@href"))

			// Numero do capitulo
			capNumber := strings.Split(htmlquery.InnerText(htmlquery.FindOne(captule, "//a/@href")), "/")[3]

			// Data de publicação
			datePublish = htmlquery.InnerText(htmlquery.FindOne(captule, "//span"))

			// Acessando a pagina dos capitulos a cada loop
			captulePage, err := htmlquery.LoadURL(linkCaptule)
			if err != nil {
				panic(err)
			}

			imagesPerPages, err := htmlquery.QueryAll(captulePage, "//article/div/div[2]/article/div[6]/div[6]")
			if err != nil {
				panic(err)
			}

			// Pegando as imagens das paginas
			for _, capImages := range imagesPerPages {
				getPagesImages, _ := htmlquery.QueryAll(capImages, "//img/@src")
				wrapperImages := []string{}
				for _, img := range getPagesImages {
					image := "https://goldenmangas.top" + htmlquery.InnerText(img)

					wrapperImages = append(wrapperImages, image)
				}

				/*
					Adiciona informações ao empacotador de capitulos
					O que estou adicionando
					-numero do capitulo
					-data de postagem
					-slice com todas as imagens desse capitulo
				*/
				CaptuleWrapper = append(CaptuleWrapper, SMangaCaptulesInfo{
					CaptuleNumber: capNumber,
					DatePublish:   datePublish,
					CaptulesPage:  wrapperImages,
				})
			}
		}

		BasicInfoManga = append(BasicInfoManga, SMangaInfoPageSource{
			Title:       title,
			Gender:      gender,
			Author:      author,
			Artist:      artist,
			Status:      status,
			Description: description,
			Captules:    CaptuleWrapper,
		})
	}
}
func ControllerScraper(file *string, scraperPerPage *int) {
	// inicia as funções
	get_routes(scraperPerPage)
	RenderScraper()

	// Cria uma pasta
	dir := "db_json"
	err := os.Mkdir(dir, 0755)
	if err != nil {
		panic(err)
	}

	// Transforma os dados do struct em json
	j, _ := json.Marshal(BasicInfoManga)

	// escreve tudo no arquivo json, e se ele n existir ele cria um
	errr := os.WriteFile(dir+"/"+*file+".json", j, 0755)
	if errr != nil {
		panic(j)
	}

	fmt.Printf("\nScraper Complete!\n File was saved in: %v.json\n", file)
	os.Exit(1)
}

func StartScraper(scraperPerPage int, outputFile string) {

	maxScraperPerPage := 347
	if scraperPerPage > maxScraperPerPage {
		scraperPerPage = 347
	}

	fileOutPut := outputFile

	ControllerScraper(&fileOutPut, &scraperPerPage)
}
