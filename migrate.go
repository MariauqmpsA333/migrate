package migrate

import (
	"regexp"
	"strings"
)

// isEmptyOrCommentOnly checks if the migration body contains only whitespace and SQL comments.
func isEmptyOrCommentOnly(body []byte) bool {
	content := string(body)
	// Remove single-line comments
	reSingleLine := regexp.MustCompile(`(?m)--.*$`)
	content = reSingleLine.ReplaceAllString(content, "")
	// Remove multi-line comments
	reMultiLine := regexp.MustCompile(`(?s)/\*.*?\*/`)
	content = reMultiLine.ReplaceAllString(content, "")
	// Check if remaining content is just whitespace
	return strings.TrimSpace(content) == ""
}

// Note: In the actual migration runner, this would be called before driver.Run(migration).
// If true, skip driver.Run and proceed to update the version.