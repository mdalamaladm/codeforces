package main

import (
  "bufio"
  "fmt"
  "os"
)

func uppercase(ch byte) string {
  code := int(ch)
  
  if code > 96 {
    return string(uint8(code - 32))
  } else {
    return string(ch)
  }
}

func lowercase(ch byte) string {
  code := int(ch)
  
  if code < 97 {
    return string(uint8(code + 32))
  } else {
    return string(ch)
  }
}

func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  var input string
  var output string
  var upper int
  var lower int
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
	
  // Logic here
  for i := 0; i < len(input); i++ {
    if int(input[i]) > 96 {
      lower++
    } else {
      upper++
    }
  }
  
  for i := 0; i < len(input); i++ {
    if upper > lower {
      output += uppercase(input[i])
    } else {
      output += lowercase(input[i])
    }
  }
	
  fmt.Fprintln(writer, output)
}