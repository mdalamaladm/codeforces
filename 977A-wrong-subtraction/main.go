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

  var number int
  var subtractTimes int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &number, &subtractTimes)
	
  // Logic here
  for i := 0; i < subtractTimes; i++ {
    if number % 10 == 0 {
      number /= 10
    } else {
      number--
    }
  }
	
  fmt.Fprintln(writer, number)
}