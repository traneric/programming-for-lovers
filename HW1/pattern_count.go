/*
Program: Substring Counting Problem
Input: A string pattern and a longer string text.
Output: The number of times that pattern appears as a substring of text.
(Remember to include overlaps.)
*/

package main

import "fmt"

func main() {
  pattern1 := "GCG"
  text1 := "GCGCG"

  pattern2 := "CG"
  text2 := "ACGTACGTACGT"

  pattern3 := "AAA"
  text3 := "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACT"+
  "TAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCAC"

  pattern4 := "TTT"
  text4 := "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCA"+
  "TAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCC"

  pattern5 := "ACT"
  text5 := "GGACTTACTGACGTACG"

  fmt.Println(PatternCount(pattern1, text1))
  fmt.Println(PatternCount(pattern2, text2))
  fmt.Println(PatternCount(pattern3, text3))
  fmt.Println(PatternCount(pattern4, text4))
  fmt.Println(PatternCount(pattern5, text5))
}

func PatternCount(pattern, text string) int {
  t := len(text)
  p := len(pattern)
  count := 0

  for i := 0; i < t-p+1; i++ {
    if text[i:i+p] == pattern {
      count++
    }
  }
  return count
}
