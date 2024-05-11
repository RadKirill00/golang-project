package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)



func handlerWiki(url, selector string) ([]string, error) {

	resp, err := http.Get(url)

	if err != nil { 
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)


	if err != nil {
		return nil, err
	}

	var texts []string

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})

	return texts, nil

}


	func wikiF(w http.ResponseWriter, r *http.Request) {
		wikiCite := "https://animestars.org/tags/%D1%81%D0%B8%D0%BB%D1%8C%D0%BD%D1%8B%D0%B9%20%D0%B3%D0%B5%D1%80%D0%BE%D0%B9/page/10/"
		
		res, err := handlerWiki(wikiCite, "img[src]")

		if err != nil { 
			log.Println(err)
		}

		for _, t := range res { 
			fmt.Println(t)
		}

		results, _ := json.Marshal(res)

		w.Write(results)
	
	}

func main() {
	http.HandleFunc("/anime", wikiF)
	http.ListenAndServe(":5959", nil)
}
