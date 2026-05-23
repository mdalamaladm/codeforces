package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  var total int
  
	fmt.Fscan(reader, &total)
	
  // Logic here
  for i := 0; i < total; i++ {
    var input string
    
    fmt.Fscan(reader, &input)
    
    length := len(input)
    
    if length > 10 {
      var output string
      
      output = string(input[0]) + strconv.Itoa(length - 2) + string(input[length - 1])
      
      fmt.Fprintln(writer, output)
    } else {
      fmt.Fprintln(writer, input)
    }
	}
}