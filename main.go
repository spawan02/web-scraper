package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
		fmt.Println("visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})
	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		author := div.Find(".author").Text()
		quote := div.Find(".text").Text()
		fmt.Printf("Quote: %s\nBy %s\n\n", quote, author)
	})
	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote", h.Text)
	// })
	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author", h.Text)
	// })
	c.Visit("https://quotes.toscrape.com")
}
