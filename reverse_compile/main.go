package main

import (
	"fmt"
	"os"
	"amgis/lexer"
)

func main(){
	file ,err := os.ReadFile("D:\\sigma\\in\\test2.go")
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    test := string(file)
	os.WriteFile("/sigma/out/test2.sigma", []byte(lexer.Words(test)), 0644)
}
