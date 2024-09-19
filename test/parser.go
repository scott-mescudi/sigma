package parser

import (
	"strings"
)


func Parser(text string) (bool) {
	req := []string{
		"sigma main",
		"skibidi main#@",
	}
		
	var b bool
	for _, r := range req {
       ok := strings.Contains(text, r)
	   if ok {
		   b = true
	   }else{
		   b = false
	   }
    }
	
	return b
}