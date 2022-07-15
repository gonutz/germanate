package main

import (
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	s, err := clipboard.ReadAll()
	if err == nil {
		s = strings.Replace(s, "aE", "ä", -1)
		s = strings.Replace(s, "AE", "Ä", -1)
		s = strings.Replace(s, "oE", "ö", -1)
		s = strings.Replace(s, "OE", "Ö", -1)
		s = strings.Replace(s, "uE", "ü", -1)
		s = strings.Replace(s, "UE", "Ü", -1)
		s = strings.Replace(s, "sS", "ß", -1)
		clipboard.WriteAll(s)
	}
}
