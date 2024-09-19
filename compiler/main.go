package main

import (
	"fmt"
	"os"
	"sigma_compiler/lexer"
	"sigma_compiler/parser"
)
var Req = []string{
		"sigma main",
		"skibidi main#@",
}

func main(){
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])



	file ,err := os.ReadFile("D:\\sigma\\tests\\out\\test2.sigma")
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    text := string(file)
	ok := parser.Parser(text, Req)
	if ok{
		compiledtext := lexer.Words(text)
		os.WriteFile("/sigma/tests/out/test2.go", []byte(compiledtext), 0644)
	}else{
		fmt.Println("error missing: 'sigma main' or 'skibidi main#@")
	}

}

//sigma build main.sigma