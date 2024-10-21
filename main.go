package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/qedhtp/von/pkg"

	"github.com/PuerkitoBio/goquery"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)




func get_query(query *string) {

	query_encode := url.QueryEscape(*query)
	// Base url   sound:https://dict.youdao.com/dictvoice?audio=cs&type=2
	//
	if len(query_encode) > 0 {
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
		pkg.GetTranslate(doc)

		
		// call the pronounce function
		// use pointer
		pkg.Pronounce(&request_url_pronounce)

	}
	
	
	
	


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
			Prompt:           promptColor("[von]>>> "),
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
			case ":exit":
				fmt.Println("Goodbye!")
				return
			case ":clear":
				clearScreen()
			default:
				get_query(&input)
			}

		}


	} else {
			// Check if the user provided a argument
		if len(os.Args) != 2 {
			color.Blue("Prints %s in blue.", "text")
			fmt.Println("Usage: von <argument>")
			return
		} else{
			// get command-line argument
			query := os.Args[1]
			get_query(&query)
	
	
		
		}
		
	}

}
