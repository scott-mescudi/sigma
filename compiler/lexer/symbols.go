package lexer

import (
	s "sigma_compiler/syntax"
)

func Symbols(test string) string {
	runes := []rune(test)
	
	for i := 0; i < len(runes); i++ {
		if val, exists := s.SigmalangMap[string(runes[i])]; exists {
			runes[i] = rune(val[0])
		}
	}
	
	return string(runes)
}
