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
  fmt.Fscan(reader, &input)
  
  fmt.Println(input)
  	
  var output string
  
  if input > 3 && input % 2 == 0 {
    output = "YES"
  } else {
    output = "NO"
  }
  
  fmt.Fprintln(writer, output)
}