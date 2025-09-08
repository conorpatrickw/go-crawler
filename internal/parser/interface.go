package parser

import "golang.org/x/net/html"

// HTMLParser defines the interface for HTML parsing operations
type HTMLParser interface {
	ExtractLinks(node *html.Node, baseUrl string) ([]string, error)
}

