package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)


func getTranslate(doc *goquery.Document) {

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
		c.Println("\n发音:")
		if pronounce_element.Length() > 0 {
			green.Printf("    英: %s    ", English)
			green.Printf("美: %s    \n", American)
		}

		// word meaning
		c.Println("\n释义:")
		for i := 0; i < word_element.Length(); i++ {
			text := word_element.Eq(i).Text()
			green.Printf("    %s", text)
		}
		fmt.Printf("\n")
		
		// word form
		if word_form_ch.Length() > 1 {
			c.Println("\n形态:")
			for i :=0; i < word_form_ch.Length(); i++ {
				text_ch := word_form_ch.Eq(i).Text()
				text_en := word_form_en.Eq(i).Text()
				
				green.Printf("    %s", text_ch)
				green.Printf(": %s", text_en)
			}
			fmt.Printf("\n")
		}

	
		// phrase
		c.Println("\n短语:")
		for i := 0; i < phrase_en.Length(); i++ {
			text_en := phrase_en.Eq(i).Text()
			text_ch := phrase_ch.Eq(i).Text()
			green.Printf("    %d.%s", i+1,text_en)
			green.Printf("  %s\n", text_ch)
		}
		fmt.Printf("\n")

		// example
		c.Println("例句:")
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

func pronounce(pronounce_url *string) {
	// mate the GET pronounce request
	response_pronounce, err := http.Get(*pronounce_url)
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

func get_query(query *string) {

	query_encode := url.QueryEscape(*query)
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

	// Call the function with *goquery.document
	getTranslate(doc)

	
	// call the pronounce function
	// use pointer
	pronounce(&request_url_pronounce)

}

// clear screen
func clearScreen()  {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}




func main() {
	// interactive model
	interactive := flag.Bool("i", false,"interactive model")

	// prompt color
	promptColor := color.New(color.FgYellow).SprintFunc()

	// Parse the command-line flags
	flag.Parse()

	// Check if the interactive flag was provided
	if *interactive {
		l, err := readline.NewEx(&readline.Config{
			Prompt:           promptColor("========> "),
			InterruptPrompt:  "^c",
			EOFPrompt:        "exit",
			HistoryFile:      "/tmp/readline.tmp",

		})
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()

		// Main interactive loop
		for {
			input, err := l.Readline()
			if err == readline.ErrInterrupt {
				// fmt.Printf("\n")
				continue
			}	
			// } else if err == io.EOF {
			// 	fmt.Println("Goodbye!")
			// 	return
			// }

			// Process user input
			switch input {
			case "q":
				fmt.Println("Goodbye!")
				return
			case "c":
				clearScreen()
			default:
				get_query(&input)
			}

		}


	} else {
			// Check if the user provided a argument
		if len(os.Args) != 2 {
			color.Blue("Prints %s in blue.", "text")
			fmt.Println("Usage: chtoen <argument>")
			return
		} else{
			// get command-line argument
			query := os.Args[1]
			get_query(&query)
	
	
		
		}
		
	}

}
