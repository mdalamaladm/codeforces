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
  
  var total int
  var target int
  var targetScore int
  output := 0
  
  // Scan every each newline/space
  fmt.Fscan(reader, &total)
  fmt.Fscan(reader, &target)
  
  // Make it zero-index number (5 => 4)
  target--
  
  // Logic here
  for i := 0; i < total; i++ {
    var score int
    
    fmt.Fscan(reader, &score)
    
    if i < target && score > 0 {
      output++
    } else if i == target {
      if score < 1 {
        break
      }
      
      output++
      
      targetScore = score
    } else if i > target && score >= targetScore {
      output++
    } else {
      break
    }
  }
  
  fmt.Fprintln(writer, output)
}