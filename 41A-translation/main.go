package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  var text string
  var txet string
  output := "YES"
  
  // Scan every each newline/space
  fmt.Fscan(reader, &text, &txet)
	
  // Logic here
  if len(text) == len(txet) {
    for i := 0; i < len(text); i++ {
      if text[i] != txet[len(text) - 1 - i] {
        output = "NO"
        break
      }
    }
  } else {
    output = "NO"
  }
	
  fmt.Fprintln(writer, output)
}