package parser

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
)

func extractFileName(filename string) (string, bool) {
	re := regexp.MustCompile(`(\w+)_ToG.go`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) > 1 {
		return matches[1], true
	} else {
		return "", false
	}
}

func extractStructName(structName string) (string, bool) {
	re := regexp.MustCompile(`(\w+)_ToG`)
	matches := re.FindStringSubmatch(structName)
	if len(matches) > 1 {
		return matches[1], true
	} else {
		return "", false
	}
}

func extractTag(tag *ast.BasicLit) map[string]string {
	m := map[string]string{}
	if tag != nil && tag.Kind == token.STRING && len(tag.Value) > 2 {
		re := regexp.MustCompile(`(\w+):"([^"]+)"`)
		for _, s := range re.FindAllStringSubmatch(tag.Value, -1) {
			m[strings.TrimSpace(s[1])] = strings.TrimSpace(s[2])
		}
	}
	return m
}
