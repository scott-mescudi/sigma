package lexer

import (
	"strings"
	s "sigma_compiler/syntax"
)


func Words(text string) string {
	for k, v := range s.SigmalangMap{
		if strings.Contains(text, k) {
			result := strings.Replace(text, k, v, -1)
			text = result
		}
	}
	return text
}
