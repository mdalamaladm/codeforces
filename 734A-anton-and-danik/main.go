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
  var games string
  var output string
  
  // Scan every each newline/space
  fmt.Fscan(reader, &total)
  fmt.Fscan(reader, &games)
	
  // Logic here
  winHalfGames := total / 2
  anton := 0
  danik := 0
  
  for i := 0; i < len(games); i++ {
    game := string(games[i])
    
    if game == "A" {
      anton++
    } else if game == "D" {
      danik++
    }
  
    if anton > winHalfGames {
      output = "Anton"
      break
    } else if danik > winHalfGames {
      output = "Danik"
      break
    }
  }
  
  if anton == danik {
    output = "Friendship"
  }
	
  fmt.Fprintln(writer, output)
}