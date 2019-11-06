/*
Program: Maximum Map Value (Strings to Ints) Problem
Input: A map freq of strings to integers.
Output: The maximum value in freq.
*/

package main

import "fmt"

func main() {
  myMap := map[string]int {
    "ACT":3,
    "GTGA":6,
    "TA":2,
  }
  fmt.Println(MaxDict(myMap))
}

func MaxDict(dict map[string]int) int {
  max := 0
  firstTimeThrough := true

  for _, value := range dict {
    if firstTimeThrough || value > max {
      max = value
      firstTimeThrough = false
    }
  }
  return max
}
