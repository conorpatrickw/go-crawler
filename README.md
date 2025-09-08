# Go Web Crawler
A simple web crawler implemented in Go that extracts links from HTML pages using DOM tree traversal.

## Features

- Single-page crawling (max depth: 1)
- Extracts all links from anchor tags (`<a href="...">`)
- Resolves relative URLs to absolute URLs
- Validates URL format before crawling

### Run the crawler 

```bash
# Run directly
go run ./cmd/crawler -url https://example.com
```
### Build and run

```bash
# Build the binary
go build -o crawler ./cmd/crawler/main.go

# Run the binary
./crawler -url https://example.com
```

### Available flags

- `-url` - URL to crawl (required)


## Implementation Details

### DOM Tree Traversal

The crawler uses **Depth-First Search (DFS)** to traverse the HTML DOM tree:

1. Parse HTML into DOM tree using `html.Parse()`
2. Recursively visit each node
3. Extract `href` attributes from anchor tags (`<a>`)
4. Resolve relative URLs to absolute URLs
5. Return collected links


## Tests

### Run all unit tests 

```bash
go test ./...
```

### Run a specific test package

```bash
go test ./pkg/crawler

```
### Potential Enhancements
- **Configurable depth**: Allow multi-level crawling with depth parameter
    - **Concurrent crawling**: Use goroutines for parallel link processing
- **Duplicate handling**: Deduplication of extracted links