package main
import (
  "fmt"
  "math"
  "container/list"
)

func sumOfPrimes(primes *(list.List)) int {
  sum := 0
  for e := primes.Front(); e != nil; e = e.Next() {
    sum += e.Value.(int)
  }
  return sum
}

func buildPrimes(upperBound int) *(list.List) {
  var primes = new(list.List)

  for i := 2; i < upperBound; i++ {
    isPrime := true
    for e := primes.Front(); e != nil && float64(e.Value.(int)) <= math.Sqrt(float64(i)); e = e.Next() {
      if i % e.Value.(int) == 0 {
        isPrime = false
        break
      }
    }
    if isPrime {
      primes.PushBack(i)
    }
  }

  return primes
}

func main() {
  var primes = buildPrimes(2000000)
  fmt.Println(sumOfPrimes(primes))
}
