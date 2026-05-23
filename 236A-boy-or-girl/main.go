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

  var username string
  distinctChar := ""
  distinctCharTotal := 0
  var output string
  
  // Scan every each newline/space
  fmt.Fscan(reader, &username)
	
  // Logic here
  for i := 0; i < len(username); i++ {
    uname := string(username[i])
    if strings.Contains(distinctChar, uname) {
      continue
    } else {
      distinctChar += uname
      distinctCharTotal++
    }
  }
  
  if distinctCharTotal % 2 == 0 {
    output = "CHAT WITH HER!"
  } else {
    output = "IGNORE HIM!"
  }
	
  fmt.Fprintln(writer, output)
}