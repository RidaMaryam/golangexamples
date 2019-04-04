package golangexamples
import (
  "fmt"
  "github.com/ehteshamz/greetings"
  )
func ConcatSlice(sliceToConcat []byte) string {
  var concatinatedString = ""
  for i := 0; i<len(sliceToConcat); i++ {
      concatinatedString = concatinatedString + string(sliceToConcat[i]) + "-"
  }
  return concatinatedString
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
      for i:=0; i<len(sliceToEncrypt); i++ {
          sliceToEncrypt[i] += byte(ceaserCount) 
          fmt.Print(string(sliceToEncrypt[i]))
      }
      fmt.Print("\n")
}

func EZGreetings(name string) string {
  return greetings.PrintGreetings(name)
}
