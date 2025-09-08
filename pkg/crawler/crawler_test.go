package crawler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/html"
)

type MockParser struct {
	ExtractLinksMock func(node *html.Node, baseUrl string) ([]string, error)
}

func (m *MockParser) ExtractLinks(node *html.Node, baseUrl string) ([]string, error) {
	return []string{"www.validtesturl.com", "www.google.com"}, nil
}

func crawlerTestClient() *crawler {
	parserMock := &MockParser{
		ExtractLinksMock: func(node *html.Node, baseUrl string) ([]string, error) {
			return []string{"www.validtesturl.com", "www.google.com"}, nil
		},
	}
	return NewClient(&CrawlerOptions{
		Parser: parserMock,
	})
}

func TestCrawlSuccess(t *testing.T) {
	crawler := crawlerTestClient()
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body><a href="www.validtesturl.com">Link1</a><a href="www.google.com">Link2</a></body></html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}))
	defer server.Close()

	links, err := crawler.Crawl(server.URL)

	if err != nil {
		t.Fatalf("Crawl error: %v", err)
	}

	expected := []string{
		"www.validtesturl.com",
		"www.google.com",
	}

	if len(links) != len(expected) {
		t.Errorf("Crawl returned %d links, want %d", len(links), len(expected))
	}

	for i, link := range links {
		if i < len(expected) && link != expected[i] {
			t.Errorf("Crawl[%d] = %s, want %s", i, link, expected[i])
		}
	}
}

func TestCrawlHTTPError(t *testing.T) {
	crawler := crawlerTestClient()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := crawler.Crawl(server.URL)
	if err == nil {
		t.Fatal("Expected error for HTTP 404, but got none")
	}
}
