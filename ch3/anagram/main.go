package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args[1:]) != 2 {
    fmt.Println("Please provide two words to compare")
  }

  firstWord := os.Args[1]
  secondWord := os.Args[2]

  firstCount := map[rune]int{}
  secondCount := map[rune]int{}

  for _, char := range(firstWord) {
    firstCount[char] += 1
  }

  for _, char := range(secondWord) {
    secondCount[char] += 1
  }

  for k, v1 := range(firstCount) {
    v2, ok := secondCount[k]

    if !ok || v1 != v2 {
      fmt.Println(firstWord, " is not an anagram of ",  secondWord)
      return
    }
  }

  fmt.Println("Is anagram")
}
