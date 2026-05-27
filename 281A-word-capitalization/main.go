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
  output := ""
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
	
  // Logic here
  for i := 0; i < len(input); i++ {
    if i == 0 {
     if int(input[i]) > 96 {
      output += string(input[i] - uint8(32))
     } else {
       output += input
       break
     }
    } else {
      output += string(input[i])
    }
  }
	
  fmt.Fprintln(writer, output)
}