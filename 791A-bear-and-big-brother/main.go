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

  var a int
  var b int
  
  fmt.Fscan(reader, &a)
  fmt.Fscan(reader, &b)
  
  year := 0
	
  // Logic here
  for a <= b {
    year++
    a *= 3
    b *= 2
  }
	
  fmt.Fprintln(writer, year)
}