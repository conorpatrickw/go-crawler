package resolver

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type resolver struct{}

func NewClient() *resolver {
	return &resolver{}
}

var urlRegex = regexp.MustCompile(`^(http|https|ftp)://[a-z0-9-]+(\.[a-z0-9-]+)+([/?].*)?$`)

func (r *resolver) ValidateURL(rawUrl string) error {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return err
	}

	if !urlRegex.MatchString(rawUrl) {
		return errors.New("URL does not match valid scheme and structure")
	}
	return nil
}

func (r *resolver) ResolveURL(href, baseresolver string) (string, error) {
	base, err := url.Parse(baseresolver)
	if err != nil {
		return "", fmt.Errorf("invalid base resolver %s: %w", baseresolver, err)
	}

	link, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("invalid href %s: %w", href, err)
	}

	resolved := base.ResolveReference(link)
	if strings.HasPrefix(resolved.String(), "http") {
		return resolved.String(), nil
	}

	return "", nil
}
