package crawler

import (
	"fmt"
	"net/http"

	"github.com/conorpatrickw/webcrawler/internal/parser"
	"golang.org/x/net/html"
)

type crawler struct {
	parser parser.HTMLParser
}

type CrawlerOptions struct {
	Parser parser.HTMLParser
}

func NewClient(options *CrawlerOptions) *crawler {
	return &crawler{
		parser: options.Parser,
	}
}

func (c *crawler) Crawl(startUrl string) ([]string, error) {
	return c.crawl(startUrl, 0)
}

func (c *crawler) crawl(currentUrl string, depth int) ([]string, error) {
	fmt.Printf("Crawling: %s \n", currentUrl)

	resp, err := http.Get(currentUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", currentUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d for URL %s", resp.StatusCode, currentUrl)
	}

	documentTree, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML from %s: %w", currentUrl, err)
	}
	links, err := c.parser.ExtractLinks(documentTree, currentUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to extract links from %s: %w", currentUrl, err)
	}

	return links, nil
}
