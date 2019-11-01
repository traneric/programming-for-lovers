/*
Program: Reverse Complement Problem
Input: A DNA string pattern.
Output: The reverse complement of pattern.
*/

package main

import "fmt"

func main() {
  dna1 := "AAAACCCGGT"
  dna2 := "ACACAC"
  dna3 := "GCTCAGCCACAACACGAGGGATACTATTATCACGGTCAGTACAACAACGCATTTGTGATCAGCAAC"+
  "GCACTAAGCTTGCCCAGGGTAGAACACGAGACGCACTCT"

  fmt.Println(ReverseComplement(dna1))
  fmt.Println(ReverseComplement(dna2))
  fmt.Println(ReverseComplement(dna3))
}

// Reverse takes a DNA string and reverses its symbols to produce a new string.
func Reverse(dna string) string {
  n := len(dna)
  reverseString := ""

  for i := (n-1); i >= 0; i-- {
    reverseString += string(dna[i])
  }
  return reverseString
}

// Complement takes a DNA string and returns its complement.
func Complement(dna string) string {
  n := len(dna)
  dna2 := ""

  for i := 0; i < n; i++ {
    switch dna[i] {
    case 'A':
      dna2 += "T"
    case 'T':
      dna2 += "A"
    case 'G':
      dna2 += "C"
    case 'C':
      dna2 += "G"
    }
  }
  return dna2
}

// ReverseComplement reverses a DNA string and returns its complement.
func ReverseComplement(pattern string) string {
  return Complement(Reverse(pattern))
}
