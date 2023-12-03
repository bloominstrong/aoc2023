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
	prevLine=currentLine
	currentLine=nextLine
	nextLine=""
	sum = sum + mySum(prevLine,currentLine,nextLine,1)
	file.Close()
	fmt.Println("Version 1 is:",sum)


	file, _ = os.Open("../input/3")
	scanner2 := bufio.NewScanner(file)
	sum = 0
	prevLine = ""
	currentLine = ""
	nextLine = string(scanner2.Bytes())
	for scanner2.Scan() {	
		prevLine = currentLine
		currentLine = nextLine
		nextLine = string(scanner2.Bytes())
		//sum = sum + mySum(prevLine,currentLine,nextLine,2)
		//fmt.Println(sum)
	}
	fmt.Println("Version 2 is:",sum)
}

func mySum (prevLine string, currentLine string, nextLine string ,version int) int {

	//prevItems := strings.Split(prevLine,".")
	currentItems := strings.Split(currentLine,".")
	//nextItems := strings.Split(nextLine,".")
	lineSum := 0
	charList := "!@#$%^&*()-_=+ /,;:'[]{}|<>?"
	for _, data := range currentItems {
		if (strings.ContainsAny(data,"0123456789")) {
			//data is a number
			startPos := strings.Index(currentLine,data)-1
			if (unicode.IsPunct(rune(data[0]))) || (startPos == -1){
				startPos++
			}
			endPos := startPos+len(data)+1
			if (unicode.IsPunct(rune(data[len(data)-1]))) || (endPos == len(currentLine)) {
				endPos--
			}
			fmt.Println("Data:",data,"\nPL:",prevLine,"\nCL:",currentLine,"\nNL:",nextLine,"Start:",startPos,"End:",endPos)
			for i := startPos; i <= endPos; i++ {
				if (endPos < len(prevLine)) {
					//fmt.Println("Pi:",i,prevLine[i])
					if (strings.ContainsAny(string(prevLine[i]),charList)) {
						dataAsNum, _ := strconv.Atoi(data)
						fmt.Println("Found Prev:",data)
						lineSum = lineSum + dataAsNum
					}
				}
				if (endPos < len(nextLine)) {	
					if (strings.ContainsAny(string(nextLine[i]),charList)) {
						dataAsNum, _ := strconv.Atoi(data) 
						fmt.Println("Found Next:",data)
						lineSum = lineSum + dataAsNum
					}
				}
			}
			if (strings.ContainsAny(data,charList)) {
				dataAsNum, _ := strconv.Atoi(strings.Trim(data,charList))
			
				fmt.Println("Found Cur:",dataAsNum)
				lineSum = lineSum + dataAsNum
			}
		}
	}

	if (version == 2){
	}

	return lineSum
}

