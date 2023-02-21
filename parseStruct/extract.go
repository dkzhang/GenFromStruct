package parseStruct

import (
	"regexp"
	"strings"
)

func extractStructName(structName string) (string, bool) {
	re := regexp.MustCompile(`(\w+)_ToG`)
	matches := re.FindStringSubmatch(structName)
	if len(matches) > 1 {
		return matches[1], true
	} else {
		return "", false
	}
}

func extractTag(tag string) map[string]string {
	m := map[string]string{}
	if tag != "" {
		re := regexp.MustCompile(`(\w+):"([^"]+)"`)
		for _, s := range re.FindAllStringSubmatch(tag, -1) {
			m[strings.TrimSpace(s[1])] = strings.TrimSpace(s[2])
		}
	}
	return m
}
