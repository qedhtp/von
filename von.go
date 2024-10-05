package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
)



func getTranslate (doc *goquery.Document) {

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


	// Determine if a word or a sentence   
	if word_element.Length() > 0 {

		// pronoounce
		fmt.Printf("\n    --------------发音---------------")
		if pronounce_element.Length() > 0 {
			fmt.Printf("\n    英: %s    ", English)
			fmt.Printf("美: %s    \n", American)
		}

		// word meaning
		fmt.Printf("\n    --------------释义---------------")
		for i := 0; i < word_element.Length(); i++ {
			text := word_element.Eq(i).Text()
			fmt.Printf("\n    %s", text)
		}
		fmt.Printf("\n")
		
		// word form
		if word_form_ch.Length() > 1 {
			fmt.Printf("\n    --------------形态---------------\n")
			for i :=0; i < word_form_ch.Length(); i++ {
				text_ch := word_form_ch.Eq(i).Text()
				text_en := word_form_en.Eq(i).Text()
				
				fmt.Printf("    %s", text_ch)
				fmt.Printf(": %s", text_en)
			}
			fmt.Printf("\n")
		}

	
		// phrase
		fmt.Printf("\n    --------------短语----------------\n")
		for i := 0; i < phrase_en.Length(); i++ {
			text_en := phrase_en.Eq(i).Text()
			text_ch := phrase_ch.Eq(i).Text()
			fmt.Printf("    %d.%s", i+1,text_en)
			fmt.Printf("  %s\n", text_ch)
		}
		fmt.Printf("\n")

		// example
		fmt.Printf("    --------------例句---------------\n")
		for i :=0; i < example_en.Length(); i++ {
			text_en := example_en.Eq(i).Text()
			text_ch := example_ch.Eq(i).Text()
			fmt.Printf("    %d.%s\n", i+1, text_en)
			fmt.Printf("      %s\n", text_ch)
		}


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
		request_url_word := "https://dict.youdao.com/result?word=" + query_encode + "&lang=en"
		request_url_pronounce := "https://dict.youdao.com/dictvoice?audio=" + query_encode +"&type=2"
		
		
		// make the GET translate request
		response_translate, err := http.Get(request_url_word)
		if err != nil {
			log.Fatal(err)
		}
		// will be closed once the main function exits
		defer response_translate.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response_translate.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Call a function with *goquery.document
		getTranslate(doc)



		// mate the GET pronounce request
		response_pronounce, err := http.Get(request_url_pronounce)
		if err != nil {
			log.Fatal(err)
		}
		// will be closed once the main function exits
		defer response_pronounce.Body.Close()

		// Create a file to save the voice binary file
		voice_file, err := os.Create("voice_tmp.mp3")
		if err != nil {
			log.Fatal(err)
		}
		defer voice_file.Close()

		// Copy the response body to the file and also a variable
		body_pronounce, err := io.ReadAll(response_pronounce.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Write the response body to the file
		_, err = voice_file.Write(body_pronounce)
		if err != nil {
			log.Fatal(err)
		}
		
		err = exec.Command("mpg123","voice_tmp.mp3","-q").Run()
		if err != nil {
			log.Fatal(err)
		}
		
		err = os.Remove("voice_tmp.mp3")
		if err != nil {
			log.Fatal(err)
		}

		
	}



}


// err := cmd.Run()

// package main

// import (
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	// Execute the 'ls' command
// 	out, err := exec.Command("ls").Output()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// Print the output
// 	fmt.Println(string(out))
// }