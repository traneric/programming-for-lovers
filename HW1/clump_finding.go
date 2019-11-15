/*
Program: Clump Finding Problem
Input:  A string text, and integers k, L, and t.
Output: All distinct k-mers forming (L, t)-clumps in text.
*/

package main

import "fmt"

func main() {
  text1 := "CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAG" +
  "GAAACATTGTAA"
  text2 := "AAAACGTCGAAAAA"
  text3 :=  "ACGTACGT"


  output1 := ClumpFinding(text1, 5, 50, 4)
  output2 := ClumpFinding(text2, 2, 4, 2)
  output3 := ClumpFinding(text3, 1, 5, 2)

  fmt.Println(output1)
  fmt.Println(output2)
  fmt.Println(output3)
}

// Returns an array of k-mers of length L, occuring at least t times.
func ClumpFinding(genome string, k, L, t int) []string {
  patterns := make([]string, 0) // Make slice of strings of length 0.
  n := len(genome) // Length of genome string.

  for i := 0; i <= n-L; i++ {
    window := genome[i:i+L] // Window range.
    freqMap := FrequencyMap(window, k) // Produce freqMap from window.

    // Iterate through freqMap. Append k-mers from clump of length L that
    // occurs t times.
    for key, val := range freqMap {
      if val >= t && inArray(patterns, key) == false {
        patterns = append(patterns, key)
      }
    }
  }
  return patterns
}

// Returns true if key is in the array.
func inArray(slice []string, key string) bool {
  for _, item := range slice {
    if item == key {
      return true
    }
  }
  return false
}

// Returns a map of all k-mers of length k.
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
