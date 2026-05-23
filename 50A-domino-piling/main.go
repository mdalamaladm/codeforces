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

  var width int
  var length int
  var output int
  
  // Scan every each newline/space
	fmt.Fscan(reader, &width)
	fmt.Fscan(reader, &length)
	
	// Logic here
	area := width * length
	output = (area - (area % 2)) / 2
	
	fmt.Fprintln(writer, output)
}