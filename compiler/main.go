package main

import (
	"fmt"
	"os"
	"sigma_compiler/lexer"
)

func main(){
	file ,err := os.ReadFile("D:\\sigma\\test\\main.sigma")
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    test := string(file)
	fmt.Println(lexer.Symbols(test))
}
