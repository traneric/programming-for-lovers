/*
Program: Frequency Map Problem
Input: A string text and an integer k.
Output: The "frequency map" of all k-mers appearing as substrings of text,
as a map of strings to integers.
*/

package main

import "fmt"

func main() {
  fmt.Println(FrequencyMap("ATATA", 3))
  fmt.Println(FrequencyMap("mamaliga", 2))
}

func FrequencyMap(text string, k int) map[string]int {
  freqMap := make(map[string]int)

  for i := 0; i < len(text)-k+1; i++ {
    pattern := text[i:i+k]

    // Shortcut: If pattern does not exist, Go creates a devault value of 0
    // and adds 1 to it. Otherwise, it adds 1 to the existing value.
    freqMap[pattern]++
  }
  return freqMap
}
