package helpers

import "regexp"

func FindPattern(content, pattern string) ([]string, error) {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(content, -1)

	return matches, nil
}
