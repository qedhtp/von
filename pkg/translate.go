package pkg

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func GetTranslate(doc *goquery.Document) {

	// pronounce
	pronounce_element := doc.Find("div.per-phone > span.phonetic")
	English := pronounce_element.First().Text()
	American := pronounce_element.Last().Text()

	// translate a word
	word_element := doc.Find("ul.basic > li.word-exp")
	word_form_ch := doc.Find("ul.word-wfs-less > li.word-wfs-cell-less > p.grey")
	word_form_en := doc.Find("ul.word-wfs-less > li.word-wfs-cell-less > span.transformation")
	phrase_en := doc.Find("div.webPhrase > ul > li > div.col2 > a.point")
	phrase_ch := doc.Find("div.webPhrase > ul > li > div.col2 > p.sen-phrase")

	// example
	example_en := doc.Find("div.col2 > div.word-exp > div.sen-eng")
	example_ch := doc.Find("div.col2 > div.word-exp > div.sen-ch")

	
	// translate a sentence
	translate_content := doc.Find("p.trans-content")


	// color 
	c := color.New(color.FgCyan, color.Underline) //.Add(color.Underline).
	green := color.New(color.FgGreen)


	// Determine if a word or a sentence   
	if word_element.Length() > 0 {

		// pronoounce
		c.Println("\nPhonetic symbol:")
		if pronounce_element.Length() > 0 {
			green.Printf("    UK: %s    ", English)
			green.Printf("US: %s    \n", American)
		}

		// word meaning
		c.Println("\nInterpretation:")
		for i := 0; i < word_element.Length(); i++ {
			text := word_element.Eq(i).Text()
			green.Printf("    %s", text)
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
		
		// word form
		if word_form_ch.Length() > 1 {
			c.Println("\nForm:")
			for i :=0; i < word_form_ch.Length(); i++ {
				text_ch := word_form_ch.Eq(i).Text()
				text_en := word_form_en.Eq(i).Text()
				
				green.Printf("    %s", text_ch)
				green.Printf(": %s", text_en)
			}
			fmt.Printf("\n")
		}

	
		// phrase
		c.Println("\nPhrase:")
		for i := 0; i < phrase_en.Length(); i++ {
			text_en := phrase_en.Eq(i).Text()
			text_ch := phrase_ch.Eq(i).Text()
			green.Printf("    %d.%s", i+1,text_en)
			green.Printf("  %s\n", text_ch)
		}
		fmt.Printf("\n")

		// example
		c.Println("Examples:")
		for i :=0; i < example_en.Length(); i++ {
			text_en := example_en.Eq(i).Text()
			text_ch := example_ch.Eq(i).Text()
			green.Printf("    %d.%s\n", i+1, text_en)
			green.Printf("      %s\n", text_ch)
		}


	} else{
		green.Printf("\n    %s\n\n", translate_content.Text())
	}
	
}