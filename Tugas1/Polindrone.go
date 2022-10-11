package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println(checkPolindrone("malam"))
	fmt.Println(checkPolindrone("siang"))
	fmt.Println(checkPolindrone("Kasur ini rusak"))
}

func checkPolindrone(input string) string {
	newInput := strings.ToLower(input)
	chars := []rune(newInput)
    var result []rune
    for i := len(chars) - 1; i >= 0; i-- {
        result = append(result, chars[i])
    }
   if newInput == string(result) {
	return "Is Polindrone true"
   }
   return "Is Polindrone false"
}