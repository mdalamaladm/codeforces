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

  var total int
  var maxHeight int
  var output int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &total, &maxHeight)
	
  // Logic here
  for i := 0; i < total; i++ {
    var height int
    fmt.Fscan(reader, &height)
    
    if (height > maxHeight) {
      output += 2
    } else {
      output += 1
    }
  }
	
  fmt.Fprintln(writer, output)
}