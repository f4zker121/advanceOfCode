package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  files, _ := os.Open("data.txt")
  counter := 0
  defer files.Close()

  strScan := bufio.NewScanner(files)

  for strScan.Scan() {
    rgbCon := map[string]int{
      "red":   0,
      "green": 0, 
      "blue":  0,
    }

    strEbal := string(strScan.Text())
    strBig := strings.Split(strEbal, " ")
    for i := 1; i < len(strBig); i++ {
      test, _ := strings.CutSuffix(strBig[i], ";")
      test, _ = strings.CutSuffix(test, ",")
      val, _ := strconv.Atoi(strBig[i-1])
      if val == 0 {
        continue
      }
      if rgbCon[test] < val {
        rgbCon[test] = val
      }
    }
    counter += rgbCon["red"] * rgbCon["green"] * rgbCon["blue"]
  }
  fmt.Println(counter)
}