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

  var price int
  var money int
  var total int
  var output int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &price, &money, &total)
	
  // Logic here
  for i := 1; i <= total; i++ {
    money -= price * i
  }
  
  if money < 0 {
    output = -money
  } else {
    output = 0
  }
	
  fmt.Fprintln(writer, output)
}