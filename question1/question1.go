package question1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(inpFile string) int {
	var line, comboString string
	var num int
	total := 0
	inp, err := os.Open(inpFile)
	if err != nil {
		panic(err)
	}
	defer inp.Close()
	re := regexp.MustCompile(`\d{1}`)
	textScan := bufio.NewScanner(inp)
	for textScan.Scan() {
		line = textScan.Text()
		digitArray := re.FindAllString(line, -1)

		if len(digitArray) == 1 {
			comboString = strings.Repeat(digitArray[0], 2)
			num, err = strconv.Atoi(comboString)
			if err != nil {
				panic(err)
			}
		} else if len(digitArray) == 0 {
			fmt.Print("No digits in array, moving on...")
			continue
		} else {
			lastIndex := len(digitArray) - 1
			comboString = digitArray[0] + digitArray[lastIndex]
			num, err = strconv.Atoi(comboString)
			if err != nil {
				panic(err)
			}
		}

		total = total + num
	}
	return total
}
