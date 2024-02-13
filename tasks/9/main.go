package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Функция для загрузки веб-страницы и сохранения ее на диск
func savePage(resp io.Reader, URL string, outputDir string) error {
	// Создаем выходной файл на диске для сохранения содержимого страницы
	fName := filepath.Base(URL) + ".html"
	fPath := filepath.Join(outputDir, fName)
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp)
	if err != nil {
		return err
	}

	fmt.Printf("Страница '%s' успешно скачана и сохранена в файл '%s'\n", URL, fPath)
	return nil
}

// Функция для парсинга ссылок на странице
func parseLinks(resp io.Reader) ([]string, error) {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(resp)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			return nil, tokenizer.Err()
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := attr.Val
					if strings.HasPrefix(link, "http") {
						links = append(links, link)
					}
					break
				}
			}
		}
	}

	return links, nil
}

func downloadSite(URL string, outputDir string) ([]string, error) {
	_, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return nil, err
		}
	}

	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = savePage(strings.NewReader(string(body)), URL, outputDir)
	if err != nil {
		return nil, err
	}

	links, err := parseLinks(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	return links, nil
}

func main() {
	URL := flag.String("url", "", "URL-адрес сайта для скачивания")
	outputDir := flag.String("output", "out", "Директория для сохранения скачанных файлов")
	flag.Parse()

	if *URL == "" {
		fmt.Println("Необходимо указать URL-адрес сайта для скачивания")
		return
	}

	links, err := downloadSite(*URL, *outputDir)
	if err != nil {
		fmt.Println("Ошибка при скачивании сайта:", err)
		return
	}

	linksLen := len(links)
	var wg sync.WaitGroup
	wg.Add(linksLen)
	for _, link := range links {
		go func(link, outputDir string) {
			defer wg.Done()
			err, _ := downloadSite(link, filepath.Join(outputDir, "inner"))
			if err != nil {
				fmt.Printf("Ошибка при скачивании ссылки %s: \n", link)
			}
		}(link, *outputDir)
	}
	wg.Wait()
}
