/*
Program: Frequency Words Problem
Input: A string text and an integer k.
Output: All most frequent k-mers in text
*/

package main

import "fmt"

func main() {
  fmt.Println(FrequentWords("ACGTTGCATGTCGCATGATGCATGAGAGCT", 4))
}

// Returns array of k-mers with the maximum occurrences (values).
func FrequentWords(text string, k int) []string {
  freqPatterns := make([]string, 0) // make slice of string of length 0
  freqMap := FrequencyMap(text, k) // generate map of kmers.
  max := MaxDict(freqMap) // finds max value of freqMap.

  for key, val := range freqMap {
    if val == max {
      freqPatterns = append(freqPatterns, key)
    }
  }
  return freqPatterns
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

// Returns max value in map.
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
