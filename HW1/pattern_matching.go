/*
Program: Pattern Matching Problem
Input: Two strings, pattern and text.
Output: All starting positions where pattern appears as a substring of text.
*/

package main

import "fmt"

func main() {
  pattern1 := "ATAT"
  text1 := "GATATATGCATATACTT"

  pattern2 := "ACAC"
  text2 := "TTTTACACTTTTTTGTGTAAAAA"

  pattern3 := "AAA"
  text3 := "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTG" +
  "ACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTA"

  pattern4 := "TTT"
  text4 :=  "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAG" +
  "TGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCC"

  pattern5 := "ATA"
  text5 := "ATATATA"

  fmt.Println(PatternMatching(pattern1, text1))
  fmt.Println(PatternMatching(pattern2, text2))
  fmt.Println(PatternMatching(pattern3, text3))
  fmt.Println(PatternMatching(pattern4, text4))
  fmt.Println(PatternMatching(pattern5, text5))
}

func PatternMatching(pattern, text string) []int {
  t := len(text)
  p := len(pattern)
  positions := make([]int, 0)

  for i := 0; i < t-p+1; i++ {
    if text[i:i+p] == pattern {
      positions = append(positions, i)
    }
  }
  return positions
}
