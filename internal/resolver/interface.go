package resolver

// URLResolver defines the interface for URL operations
type URLResolver interface {
	ResolveURL(href, baseUrl string) (string, error)
	ValidateURL(rawUrl string) error
}