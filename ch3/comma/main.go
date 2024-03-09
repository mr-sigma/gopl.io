// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
        "bytes"
        "strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
  var buff bytes.Buffer
  var str string

  if strings.HasPrefix(s, "-") {
    buff.WriteByte('-')
    str = s[1:]
  } else {
    str = s
  }

  n := len(str)
  if n <= 3 {
    return s
  }

  dot := strings.LastIndex(str, ".")

  if dot == -1 {
    dot = n
  }


  for i := 0; i < n; i++ {
    // 012345
    // 1000
    // 1,000
    // 15000
    // 100000
    // 123,456,789
    if i > 0 && i < dot && (dot - i) % 3 == 0 {
      buff.WriteByte(',')
    }
    buff.WriteByte(str[i])
  }

  return buff.String()
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}
// 	return comma(s[:n-3]) + "," + s[n-3:]
// }

//!-
