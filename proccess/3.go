package main

	import (
		"fmt"
		"strings"
		"strconv"
		"os"
		"bufio"
		"unicode"
	)

func main() {
	//from input
	file, _ := os.Open("../input/3")
	scanner := bufio.NewScanner(file)
	sum := 0
	prevLine := ""
	currentLine := ""
	nextLine := string(scanner.Bytes())
	for scanner.Scan() {
		prevLine = currentLine
		currentLine = nextLine
		nextLine = string(scanner.Bytes())
		sum = sum + mySum(prevLine,currentLine,nextLine,1)
		//fmt.Println(sum)
	}
	file.Close()
	fmt.Println("Version 1 is:",sum)


	file, _ = os.Open("../input/3")
	scanner2 := bufio.NewScanner(file)
	sum = 0
	for scanner2.Scan() {	
		prevLine = currentLine
		currentLine = nextLine
		nextLine = string(scanner.Bytes())
		sum = sum + mySum(prevLine,currentLine,nextLine,2)
		//fmt.Println(sum)
	}
	fmt.Println("Version 2 is:",sum)
}

func mySum (prevLine string, currentLine string, nextLine string ,version int) int {

	//prevItems := strings.Split(prevLine,".")
	currentItems := strings.Split(currentLine,".")
	//nextItems := strings.Split(nextLine,".")
	lineSum := 0

	for _, data := range currentItems {
		if (strings.ContainsAny(data,"0123456789")) {
			//data is a number
			startPos := strings.Index(currentLine,data)-1
			if (unicode.IsSymbol(rune(data[0]))) || (startPos == -1){
				startPos++
			}
			endPos := startPos+len(data)+1
			if (unicode.IsSymbol(rune(data[len(data)-1]))) || (endPos == len(currentLine)) {
				endPos--
			}
			fmt.Println("Data:",data,"CurLine:",currentLine,"LineLen:",len(prevLine),"Start:",startPos,"End:",endPos)
			for i := startPos; i < endPos; i++ {
				if (endPos < len(prevLine)) {
					if (unicode.IsSymbol(rune(prevLine[i]))) {
						dataAsNum, _ := strconv.Atoi(data) 
						lineSum = lineSum + dataAsNum
					}
				}
				if (endPos < len(nextLine)) {
					if (unicode.IsSymbol(rune(nextLine[i]))) && (endPos<len(nextLine))  {
						dataAsNum, _ := strconv.Atoi(data) 
						lineSum = lineSum + dataAsNum
					}
				}
			}
			if (unicode.IsSymbol(rune(currentLine[startPos]))) || (unicode.IsSymbol(rune(currentLine[endPos]))) {
				dataAsNum, _ := strconv.Atoi(data) 
				lineSum = lineSum + dataAsNum
			}
		}
	}

	if (version == 2){
	}

	return lineSum
}

