package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
    // Sample HTML document
    htmlContent := `
    <html>
        <body>
            <div class="item">First Item</div>
            <div class="item">Second Item</div>
            <div class="item">Third Item</div>
        </body>
    </html>`

    // Load the HTML document into Goquery
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
    if err != nil {
        fmt.Println("Error loading HTML:", err)
        return
    }

    // Use Find() to select the elements
    items := doc.Find("div.item")

    // Access the text of each item using Get()
    for i := 0; i < items.Length(); i++ {
        text := items.Eq(i).Text() // Get the *html.Node at index i
        fmt.Println(text)     // Print the text
    }
}
