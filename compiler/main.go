package main

import (
	"fmt"
	"os"
	"sigma_compiler/lexer"
)

func main(){
	file ,err := os.ReadFile("D:\\sigma\\out\\test2.sigma")
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    test := string(file)
	os.WriteFile("/sigma/out/test2.go", []byte(lexer.Words(test)), 0644)
}
