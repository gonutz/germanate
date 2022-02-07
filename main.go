package main

import (
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	s, err := clipboard.ReadAll()
	if err == nil {
		s = strings.Replace(s, "ae", "ä", -1)
		s = strings.Replace(s, "Ae", "Ä", -1)
		s = strings.Replace(s, "oe", "ö", -1)
		s = strings.Replace(s, "Oe", "Ö", -1)
		s = strings.Replace(s, "ue", "ü", -1)
		s = strings.Replace(s, "Ue", "Ü", -1)
		s = strings.Replace(s, "sss", "ß", -1)
		clipboard.WriteAll(s)
	}
}
