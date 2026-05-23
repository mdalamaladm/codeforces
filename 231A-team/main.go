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
  var output int
  
	fmt.Fscan(reader, &total)
	
	// Logic here
	for i := 0; i < total; i++ {
	  var certain string
	  var totalCertain int
	 
	  for j := 0; j < 3; j++ {
	    fmt.Fscan(reader, &certain)
	    
	    // Skip if totalCertain already 2 before j = 2
	    if totalCertain > 1 {
	       fmt.Fscan(reader, &certain)
	       break
	    }
	   
	    if certain == "1" {
	      totalCertain++
	      
	      if totalCertain > 1 {
	        output++
	      }
	    }
	  }
	}
	
	fmt.Fprintln(writer, output)
}