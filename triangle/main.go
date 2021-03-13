package main

/*
What to make:
 - a right triangle of # symbols
ex)
#
##
###
####
#####

Constraints:
 - no more than 2 print statements
 - has to respond to an integer. ie) maxSize := 5

Plan/steps:
print 1
print 2
print 3
print 4
print 5

bonus:
- when you reach max size, decrement back down to 1
ex) maxSize3
#
##
###
##
#
*/

import (
  "fmt"
  "strings"
)


func buildString(maxSize int) string {
  row := ""
  for i := 0; i <= maxSize; i++ {
    row = row + "#"
  }
  return row
}

func reverseTriangle(maxSize int) {
  row := buildString(maxSize)
  for i := len(row) - 1; i > 0; i-- {
    row = strings.TrimSuffix(row, "#")
    fmt.Println(row)
  }
}

func triangle(maxSize int) {
  row := ""
  //fmt.Println(maxSize)
  for i := 0; i <= maxSize; i++ {
    if i == maxSize {
    } else {
      row = row + "#"
      fmt.Println(row)
    }
  }
}

func main() {
  //triangle(5)
  triangle(9)
  reverseTriangle(9)
}
