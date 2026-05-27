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

  var input string
  var output int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
  fmt.Fscan(reader, &input)
  
  prev := ""
	
  // Logic here
  for _, char := range input {
    charStr := string(char)
    
    if charStr == prev {
      output++
    }
    
    prev = charStr
  }
	
  fmt.Fprintln(writer, output)
}