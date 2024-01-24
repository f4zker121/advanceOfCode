package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	files, _ := os.Open("data.txt")
	defer files.Close()
	counter := 0
	myNumsArr := regexp.MustCompile("[0-9]")
	strScan := bufio.NewScanner(files)
	for strScan.Scan() {
		strReal := string(strScan.Text())
		strNums := myNumsArr.FindAllString(strReal, -1)
		a := strNums[0] + strNums[len(strNums)-1]
		i, _ := strconv.Atoi(a)
		counter += i
	}
	fmt.Println(counter)
}
