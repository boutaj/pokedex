package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmedFields := strings.Fields(text)
	for i, v := range trimmedFields {
		trimmedFields[i] = strings.ToLower(v)
	}
	return trimmedFields
}