package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
  "gopl.io/ch2/tempconv"
)

func main() {
  if len(os.Args) == 1 {
    readFromStdin()
  }

  for _, arg := range(os.Args[1:]) {
    displayArgs(arg)
  }
}

func readFromStdin() {
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Println("Enter a number: ")
    text, _ := reader.ReadString('\n')
    text = strings.Trim(text, "\n")

    displayArgs(text)
  }
}

func displayArgs(arg string) {
  num, err := strconv.ParseFloat(arg, 64)

  if err != nil {
    fmt.Fprintf(os.Stderr, "unitconv: %v", err)
    os.Exit(1)
  }

  f := tempconv.Fahrenheit(num)
  c := tempconv.Celsius(num)
  lbs := Pound(num)
  kgs := Kilogram(num)
  in := Inch(num)

  fmt.Printf("%s = %s, %s = %s, %s = %s, %s =  %s, %s = %s = %s\n\n", f, tempconv.FToC(f), c, tempconv.CToF(c), lbs, PoundToKilogram(lbs), kgs, KilogramToPound(kgs), in, InchToFoot(in), InchToMeter(in))
}

