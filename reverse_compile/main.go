package main

import (
	"fmt"
	"os"
	"amgis/lexer"
)

func main(){
	file ,err := os.ReadFile("D:\\sigma\\out\\sigma.go")
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    test := string(file)
	os.WriteFile("/sigma/out/main.sigma", []byte(lexer.Words(test)), 0644)
}
