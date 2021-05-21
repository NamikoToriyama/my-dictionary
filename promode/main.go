package promode

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var sc = bufio.NewScanner(os.Stdin)

// ProMode ... this mode is the word is searched by using actually dictionally.
func ProMode() {
	var s string
	for {
		fmt.Printf("[Pro]Search Word > ")
		if sc.Scan() {
			s = sc.Text()
		}
		if s == "" {
			fmt.Println("Bye:)")
			break
		}
		m, err := DoScraping(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf(m + "\n\n")
	}
}

// DoScraping ... print the meaning of a word by scraping.
func DoScraping(s string) (string, error) {
	path := "https://ejje.weblio.jp/content/" + s
	doc, err := goquery.NewDocument(path)
	if err != nil {
		fmt.Print("url scarapping failed")
		return "", err
	}
	// pick up word's meaing in html
	summary := doc.Find("div#contentWrp").Find("div#summary")
	meaning, err := summary.Find("div.summaryM").Find("td.content-explanation").Html()
	if meaning == "" {
		return "", errors.New(s + ": is no found")
	}
	if err != nil {
		return "", errors.New("dom get failed")
	}
	return meaning, nil
}
