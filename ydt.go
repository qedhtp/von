package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func getPronounce (doc *goquery.Document) {

	pronounce_element := doc.Find("div.per-phone > span.phonetic")

	English := pronounce_element.First().Text()
	American := pronounce_element.Last().Text()

	

	if pronounce_element.Length() > 0 {
		fmt.Printf("\n    英: %s    ", English)
		fmt.Printf("美: %s    \n", American)
	}

}

func getTranslate (doc *goquery.Document) {
	// translate a word
	translate_element := doc.Find("ul.basic > li.word-exp")
	word_form := doc.Find("ul.word-wfs-less > li.word-wfs-cell-less")
	phrase := doc.Find("div.col2 > a.point")
	
	// translate a sentence
	translate_content := doc.Find("p.trans-content")

	

	// Determine if a word or a sentence
	if translate_element.Length() > 0 {
		for i := 0; i < translate_element.Length(); i++ {
			text := translate_element.Eq(i).Text()
			fmt.Printf("\n    %s", text)
		}
		fmt.Printf("\n\n")
	
		for i :=0; i < word_form.Length(); i++ {
			text := word_form.Eq(i).Text()
			fmt.Printf("    %s", text)
		}
		fmt.Printf("\n\n")
	
		for i := 0; i < phrase.Length(); i++ {
			text := phrase.Eq(i).Text()
			fmt.Printf("    %s", text)
		}
		fmt.Printf("\n")


	} else{
		fmt.Printf("\n    %s\n\n", translate_content.Text())
	}
	
}


func main() {
	
	//pattern := 

	// Check if the user provided a argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: chtoen <argument>")
	} else{
		// get command-line argument
		query := os.Args[1]
		query_encode := url.QueryEscape(query)
		// Base url   sound:https://dict.youdao.com/dictvoice?audio=cs&type=2
		request_url := "https://dict.youdao.com/result?word=" + query_encode + "&lang=en"
		
		
		// make the GET request
		response, err := http.Get(request_url)
		if err != nil {
			log.Fatal(err)
		}
		// will be closed once the main function exits
		defer response.Body.Close()
		
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Call a function with *goquery.document
		getPronounce(doc)
		getTranslate(doc)


		
	}



}
