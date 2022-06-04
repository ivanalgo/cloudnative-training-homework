package main

import "fmt"
import "time"

func main() {
    ch := make(chan int, 10)

    go func(ch chan <- int) {
      for i := 0; true; i++ {
          ch <- i
          fmt.Println("Producer : ", i)
          time.Sleep(time.Second)
       }
    }(ch)

    for d := range ch {
      fmt.Println("Comsumer : ", d)
    }
}
