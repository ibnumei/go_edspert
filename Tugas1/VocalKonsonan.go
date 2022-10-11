package main

import (
  "fmt"
  "strings"
  "regexp"
)

func main() {
  fmt.Println(isVocal(strings.ToUpper("a"))) //Vocal
  fmt.Println(isVocal(strings.ToUpper("a1"))) //Is Not Valid
  fmt.Println(isVocal(strings.ToUpper("aa"))) //More Than 1 Character
  fmt.Println(isVocal(strings.ToUpper("b"))) //Konsonan
}

func isVocal(input string) string {
  vocalAbjad := "AIUEO"
  isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

  if !isAlpha(input) {
    return "Is Not Valid"
  }
  if len(input) > 1 {
    return "More Than 1 Character"
  }
  if strings.Contains(vocalAbjad, input) {
    return "Vocal"
  }
  return "Konsonan"
}