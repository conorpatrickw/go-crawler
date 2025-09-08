package parser

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type MockResolver struct {
	ResolveURLMock func(href, baseUrl string) (string, error)
}

func (m *MockResolver) ResolveURL(href, baseUrl string) (string, error) {
	if m.ResolveURLMock != nil {
		return m.ResolveURLMock(href, baseUrl)
	}
	return href, nil
}

func (m *MockResolver) ValidateURL(url string) error {
	return nil
}

func TestExtractLinks(t *testing.T) {
	resolver := &MockResolver{
		ResolveURLMock: func(href, baseUrl string) (string, error) {
			return "http://example.com" + href, nil
		},
	}
	parser := NewClient(&ClientOptions{Resolver: resolver})

	tests := []struct {
		name     string
		htmlStr  string
		expected []string
	}{
		{
			name:     "single link",
			htmlStr:  `<a href="/page">Link</a>`,
			expected: []string{"http://example.com/page"},
		},
		{
			name:     "multiple links",
			htmlStr:  `<a href="/page1">Link1</a><a href="/page2">Link2</a>`,
			expected: []string{"http://example.com/page1", "http://example.com/page2"},
		},
		{
			name:     "nested links",
			htmlStr:  `<div><a href="/outer">Outer</a><div><a href="/inner">Inner</a></div></div>`,
			expected: []string{"http://example.com/outer", "http://example.com/inner"},
		},
		{
			name:     "no links",
			htmlStr:  `<div>No links here</div>`,
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node, _ := html.Parse(strings.NewReader(test.htmlStr))
			links, err := parser.ExtractLinks(node, "http://example.com")

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(links) != len(test.expected) {
				t.Errorf("got %d links, want %d", len(links), len(test.expected))
			}

			for i, link := range links {
				if i < len(test.expected) && link != test.expected[i] {
					t.Errorf("link[%d] = %s, want %s", i, link, test.expected[i])
				}
			}
		})
	}
}
