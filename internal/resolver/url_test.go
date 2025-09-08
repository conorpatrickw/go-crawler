package resolver

import (
	"testing"
)

func TestValidateUrl(t *testing.T) {
	resolver := NewClient()

	testCases := []struct {
		name        string
		url         string
		shouldError bool
	}{
		{"Valid HTTP URL with golang", "http://golang.org", false},
		{"Valid HTTPS URL with golang", "https://golang.org", false},
		{"Valid URL with www.golang", "http://www.golang.org", false},
		{"Valid URL with query for golang", "http://golang.org/path?query=value", false},
		{"Invalid URL without scheme", "golang.org", true},
		{"Invalid URL with empty string", "", true},
		{"Invalid URL with only scheme", "http://", true},
		{"Invalid malformed URL", "http://golang", true},
	}

	for _, testCase := range testCases {
		error := resolver.ValidateURL(testCase.url)
		if (error != nil) != testCase.shouldError {
			t.Errorf("validateUrl() returned %v for URL %s, expected error: %v", error, testCase.url, testCase.shouldError)
		}
	}
}

func TestResolveURL(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		href     string
		baseUrl  string
		expected string
		hasError bool
	}{
		{
			name:     "relative path",
			href:     "page.html",
			baseUrl:  "http://example.com",
			expected: "http://example.com/page.html",
			hasError: false,
		},
		{
			name:     "absolute path",
			href:     "/about",
			baseUrl:  "http://example.com/home",
			expected: "http://example.com/about",
			hasError: false,
		},
		{
			name:     "external URL",
			href:     "http://other.com",
			baseUrl:  "http://example.com",
			expected: "http://other.com",
			hasError: false,
		},
		{
			name:     "parent directory",
			href:     "../page.html",
			baseUrl:  "http://example.com/dir/",
			expected: "http://example.com/page.html",
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := client.ResolveURL(test.href, test.baseUrl)

			if test.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("got %s, want %s", result, test.expected)
			}
		})
	}
}
