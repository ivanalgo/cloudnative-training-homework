package main

import (
    "fmt"
)

func main() {
  str := [5]string{"i", "am", "stupid", "and", "week"}

  for i, _ := range str {
    if str[i] == "stupid" {
       str[i] = "smart"
    } else if str[i] == "week" {
      str[i] = "strong"
    }
  }

  fmt.Println(str)
}
