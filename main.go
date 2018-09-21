package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/advancedlogic/GoOse"
)

func main() {
	url := flag.String("url", "", "URL of article to extract")
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(3)
	}

	g := goose.New()
	article, _ := g.ExtractFromURL(*url)
	fmt.Print(article.CleanedText)
}
