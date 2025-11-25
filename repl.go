package main

import "strings"

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	lowerAndTrimmed := strings.TrimSpace(lower)
	return strings.Fields(lowerAndTrimmed)
}
