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

  var input string
  var totalLucky int
  var output string
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
	
  // Logic here
  for i := 0; i < len(input); i++ {
    num, _ := strconv.Atoi(string(input[i]))
    
    if num == 4 || num == 7 {
      totalLucky++
    }
  }
  
  if totalLucky == 4 || totalLucky == 7 {
    output = "YES"
  } else {
    output = "NO"
  }
	
  fmt.Fprintln(writer, output)
}