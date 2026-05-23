package main

import (
  "bufio"
  "fmt"
  "os"
)

func AbsInt(n int) int {
  if n < 0 {
      return -n
  }
  return n
}


func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  var number int
  var output int
  
  // Logic here
  for row := 1; row <= 5; row++ {
    for col := 1; col <= 5; col++ {
      fmt.Fscan(reader, &number)
      
      if number == 1 {
        output = AbsInt(3 - row) + AbsInt(3 - col)
        
        break
      }
    }
  }

  fmt.Fprintln(writer, output)
}