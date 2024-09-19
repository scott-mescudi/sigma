package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sigma_compiler/lexer"
	"sigma_compiler/parser"
	"strings"
)
var Req = []string{
		"sigma main",
		"skibidi main#@",
}

func main(){
	if len(os.Args) != 3{
		fmt.Println("Usage: sigma <flag> <file>")
        return
	}
	flag := os.Args[1]
	filename := os.Args[2]

	switch flag {
		case "build":
            build(filename)
        default:
            fmt.Println("Invalid flag. Use 'build'")
	}
}

func buildGoFilr(filename string) error{
	ext := filepath.Ext(filename)
	if ext!= ".go"{
        return fmt.Errorf("failed to compile")
    }

	cmd := exec.Command("go" ,"build", filename)
	err := cmd.Run()
	if err!= nil {
        return fmt.Errorf("failed to compile")
    }
	
	return nil
}

func isProgramInstalled(program string) bool {
	_, err := exec.LookPath(program)
	return err == nil
}


func checkfile(filename string) (newfile string, ok bool){
	ext := filepath.Ext(filename)
	if ext!= ".sigma"{
        return "", false
    }

	_, err := os.Stat(filename)
	if os.IsNotExist(err){
        return "", false
    }

	ff := strings.TrimSuffix(filename, ".sigma")
	return ff + ".go", true
}

func build(filename string){
	if !isProgramInstalled("go"){
        fmt.Println("go compiler is not installed")
        return
    }

	newfile, ok := checkfile(filename)
	if !ok{
		fmt.Println("file does not exist")
        return
	}

	file ,err := os.ReadFile(filename)
	if err!= nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
	
    text := string(file)
	ok = parser.Parser(text, Req)
	if ok{
		compiledtext := lexer.Words(text)
		err := os.WriteFile(newfile, []byte(compiledtext), 0644)
		if err!= nil {
            fmt.Printf("Error writing file: %v\n", err)
            return
        }

		defer os.Remove(newfile)
		err = buildGoFilr(newfile)
		if err != nil{
			fmt.Println(err)
			return
        }

		fmt.Println("Compilation successful")
	}else{
		fmt.Println("error missing: 'sigma main' or 'skibidi main#@")
	}

}

