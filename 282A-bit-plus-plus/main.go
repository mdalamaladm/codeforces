package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  
  var lines int
  var output int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &lines)
  
  // Logic here
  for i := 0; i < lines; i++ {
    var statement string
    
    fmt.Fscan(reader, &statement)
   
    if strings.Contains(statement, "++") {
      output++
    } else if strings.Contains(statement, "--") {
      output--
    }
  }
  
  fmt.Fprintln(writer, output)
}