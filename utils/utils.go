package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func Search(text string, pattern string) bool {

	return strings.Contains(strings.ToLower(text), strings.ToLower(pattern))

}

func ConvertToCamelCase(text string) string {
	if len(text) <= 0 {
		return text
	}

	runes := []rune(text)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)

}

func RemoveStringSpaces(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func ForceSnakeCase(text string) string {
	return strings.ReplaceAll(strings.TrimSpace(text), " ", "_")
}

func SnakeCaseToPascalCase(text string) string {
	pascalCase := ""
	for _, v := range strings.Split(text, "_") {
		pascalCase += CapitilizeFirst(v)
	}

	return pascalCase
}

func CapitilizeFirst(text string) string {
	if len(text) <= 1 {
		return text
	}

	runes := []rune(text)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func CamelToSnake(camel string) string {
	re := regexp.MustCompile("([A-Z])")
	return re.ReplaceAllStringFunc(camel, func(s string) string {
		return "_" + strings.ToLower(s)
	})
}

func SnakeCaseToCamelCase(text string) string {
	camelCase := ""
	for i, v := range strings.Split(text, "_") {
		if i == 0 {
			camelCase += v
		} else {
			camelCase += CapitilizeFirst(v)

		}
	}

	return camelCase
}
