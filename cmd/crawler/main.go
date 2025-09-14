package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/conorpatrickw/webcrawler/internal/parser"
	"github.com/conorpatrickw/webcrawler/internal/resolver"
	"github.com/conorpatrickw/webcrawler/pkg/crawler"
)

func main() {
	url := flag.String("url", "", "starting URL to crawl")
	flag.Parse()

	if *url == "" {
		log.Fatal("Please provide a starting URL using -url flag")
	}
	resolver := resolver.NewClient()

	if err := resolver.ValidateURL(*url); err != nil {
		log.Fatal("Invalid URL provided:", err)
	}

	parserClient := parser.NewClient(&parser.ClientOptions{
		Resolver: resolver,
	})
	crawlerClient := crawler.NewClient(&crawler.CrawlerOptions{
		Parser: parserClient,
	})

	links, err := crawlerClient.Crawl(*url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d links:\n", len(links))
	for _, link := range links {
		fmt.Println(link)
	}
}
