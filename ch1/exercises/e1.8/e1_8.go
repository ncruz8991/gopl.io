// Page 17
package e1_8

import "strings"

var protocol = "http://"

// Adds the "http://" prefix to the url string if it needs it.
func ParseUrl(url string) string {
	if !strings.HasPrefix(url, protocol) {
		return protocol + url
	}
	return url
}
