package parser

import (
	"github.com/conorpatrickw/webcrawler/internal/resolver"
	"golang.org/x/net/html"
)

type parser struct {
	resolver resolver.URLResolver
}

type ClientOptions struct {
	Resolver resolver.URLResolver
}

func NewClient(options *ClientOptions) *parser {
	return &parser{
		resolver: options.Resolver,
	}
}

func (p *parser) ExtractLinks(node *html.Node, baseUrl string) ([]string, error) {
	var links []string

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				if link, err := p.resolver.ResolveURL(attr.Val, baseUrl); err != nil {
					return nil, err
				} else {
					if link != "" {
						links = append(links, link)
					}
				}
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		newNodeLinks, err := p.ExtractLinks(child, baseUrl)
		if err != nil {
			return nil, err
		}
		links = append(links, newNodeLinks...)
	}

	return links, nil
}
