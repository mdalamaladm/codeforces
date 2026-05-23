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

  var stringA string
  var stringB string
  output := 0
  
  // First string
  fmt.Fscan(reader, &stringA)
  // Second string
  fmt.Fscan(reader, &stringB)
  
  for i := 0; i < len(stringA); i++ {
    byteA := int(stringA[i])
    byteB := int(stringB[i])
    
    if byteA > 96 {
      byteA -= 32
    }
  
    if byteB > 96 {
      byteB -= 32
    }
    
    if byteA > byteB {
      output = 1
      break
    } else if byteB > byteA {
      output = -1
      break
    }
  }
	
  fmt.Fprintln(writer, output)
}