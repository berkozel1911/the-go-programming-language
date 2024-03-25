// Exercise 1.1: Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.

package main 

import (
   "fmt"
   "os"
   "time"
   "strings"
)

func main() {
   start := time.Now()
   s := ""
   for _,val := range os.Args[1:] {
      s += val
   }
   t := time.Now()
   fmt.Println(s)
   fmt.Println("Time elapsed", t.Sub(start))

   start = time.Now()
   fmt.Println(strings.Join(os.Args[1:], " "))
   t = time.Now()
   fmt.Println("Time elapsed", t.Sub(start))
}
