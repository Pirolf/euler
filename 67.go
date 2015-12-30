package main

import (
  "fmt"
  "io/ioutil"
  "math"
  "log"
  "os"
  "strings"
  "strconv"
)

func findPath(filename string, maxLines int) int {
    contents, error := ioutil.ReadFile(filename)
    if error != nil {
      log.Fatal(error)
    }

    lines := strings.Split(string(contents), "\n")

    var maxSubscores [101]int
    var intNumbers [101]int
    //maxLines-1: remove trailing empty line
    for i := maxLines-1; i >= 0; i-- {
      numbers := strings.Split(lines[i], " ")
      for j := 0; j < len(numbers); j++ {
        n := numbers[j]
        //style fix: inputs contain numbers padded with 0 in the front
        if n[0:1] == "0" && len(n) >= 2 {
          n = n[1:len(n)]
        }
        number, err := strconv.Atoi(n)
        if err != nil {
          log.Fatal(err)
        }
        intNumbers[j] = number
      }
      if i == maxLines-1 {
        maxSubscores = intNumbers
      } else {
        for k := 0; k < i+1; k++ {
          maxSubscores[k] = intNumbers[k] + int(math.Max(float64(maxSubscores[k]), float64(maxSubscores[k+1])))
        }
      }
    }
    return maxSubscores[0]
}

func main() {
  path, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(findPath(path + "/p067_triangle.txt", 100))
}
