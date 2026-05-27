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

  var input int
  var output int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
	
  // Logic here
  for input > 0 {
    if input - 5 >= 0 {
      input -= 5
      output++
    } else if input - 4 >= 0 {
      input -= 4
      output++
    } else if input - 3 >= 0 {
      input -= 3
      output++
    } else if input - 2 >= 0 {
      input -= 2
      output++
    } else if input - 1 >= 0 {
      input -= 1
      output++
    }
  }
	
  fmt.Fprintln(writer, output)
}