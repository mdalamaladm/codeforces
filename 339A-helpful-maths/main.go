package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  // Fast I/O setup
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  var input string
  output := ""
  
  // Scan every each newline/space
  fmt.Fscan(reader, &input)
  
  // Logic here
  nums := strings.Split(input, "+")
  ones := 0
  twos := 0
  threes := 0
  
  for i := 0; i < len(nums); i++ {
    num, _ := strconv.Atoi(nums[i])
    
    if num == 1 {
      ones++
    } else if num == 2 {
      twos++
    } else if num == 3 {
      threes++
    }
  }
  
  for ones > 0 || twos > 0 || threes > 0 {
    if ones > 0 {
      output += "1"
      ones--
    } else if twos > 0 {
      output += "2"
      twos--
    } else if threes > 0 {
      output += "3"
      threes--
    }
    
    output += "+"
  }
  
  output = output[:len(output) - 1]
	
  fmt.Fprintln(writer, output)
}