package parser

import (
	"strings"
)


func Parser(text string ,req []string) (bool) {		
	var b bool = true
	for _, r := range req {
		if !b {
			return false
		}

       ok := strings.Contains(text, r)
	   if ok {
		  b = true
	   }else{
		   b = false
	   }
    }
	
	return b
}