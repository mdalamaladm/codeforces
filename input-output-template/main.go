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
  var output
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
	
  // Logic here
	
  fmt.Fprintln(writer, output)
}