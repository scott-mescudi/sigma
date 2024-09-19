package main

import "fmt"

func main() {
    sigmalangMap := map[string]string{
        "intywinty":      "int",
        "Slay":           "float32",
        "bigslay":        "float64",
        "wordsnletters":  "string",
        "oneandtwo":      "bool",
        "Periodt":        "const",
        "quandaledingle": "var",
        "Yeet":           ":=",
        "rizz":           "+",
        "aura":           "-",
        "bussin":         "*",
        "ratio":          "%",
        "lowkey":         "<",
        "highkey":        ">",
        "Sus":            "&",
        "straw":          "|",
        "nuhuh":          "!",
        "Vibe":           "=",
        "<div>":          "if",
        "</div>":         "else",
        "Drip":           "for",
        "skibidi":        "func",
        "spinback":       "return",
        "uhhthething":    "type",
        "mobnke":         "struct",
        "@":              ")",
        "#":              "(",
        "livvy":          "]",
        "babygronk":      "[",
        "$":              "{",
        "%":              "}",
        "|":              ":",
        "bringin":        "import",
    }

    for k, v := range sigmalangMap{
        fmt.Printf("'%v' = '%v'\n", v,k)
    }
}
