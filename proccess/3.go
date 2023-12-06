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
	//currentItems := strings.Split(currentLine,".")
	//nextItems := strings.Split(nextLine,".")
	lineSum := 0
	charList := "!@#$%^&*()-_=+ /,;:'[]{}|<>?`~"
	currentNumbers := strings.FieldsFunc(currentLine, func(c rune) bool { return !unicode.IsNumber(c)})
	found := false
	numberOccurance := make([]int,len(currentNumbers),20)
	for i, d1 := range currentNumbers {
		for j, d2 := range currentNumbers {
			if (i <= j) && (d1 == d2) {
				numberOccurance[i]++
			}
		}
	}
	//fmt.Println("CN:",currentNumbers,"\nNO:",numberOccurance)
	for j, data := range currentNumbers {
			found = false
			startPos := 0
			if (numberOccurance[j]==2) {
				firstPos := strings.Index(currentLine,data)
				startPos = strings.Index(currentLine[j:],data) + firstPos
				//fmt.Println("DUPLICATE")
				//fmt.Println("Data:",data,"\nPL:",prevLine,"\nCL:",currentLine,"\nNL:",nextLine,"Start:",startPos)
			} else if (numberOccurance[j] > 2) {
				fmt.Println("TOO MANY DUPLICATES FOUND")
			} else {
				startPos = strings.Index(currentLine,data)-1
			}
			endPos := startPos+len(data)+1
			
			if (endPos == len(currentLine)) {
				endPos--
			}
			if (startPos == -1){
				startPos++
			}
			if (startPos != 0) && (endPos != 139) && (!unicode.IsPunct(rune(currentLine[startPos])) || !unicode.IsPunct(rune(currentLine[endPos]))) {
				fmt.Println("PROBLEM LINES\n",prevLine,"\n",currentLine,"\n",nextLine)
			}
			//fmt.Println("Data:",data,"\nPL:",prevLine,"\nCL:",currentLine,"\nNL:",nextLine,"Start:",startPos,"End:",endPos)
		PosLoop:
			for i := startPos; i <= endPos; i++ {
				if (endPos < len(prevLine)) {
					if (strings.ContainsAny(string(prevLine[i]),charList)) {
						dataAsNum, _ := strconv.Atoi(data)
						//fmt.Println("Found Prev:",data)
						lineSum = lineSum + dataAsNum
						found = true
						break PosLoop
					}
				}
				if (endPos < len(nextLine)) {	
					if (strings.ContainsAny(string(nextLine[i]),charList)) {
						dataAsNum, _ := strconv.Atoi(data) 
						//fmt.Println("Found Next:",data)
						lineSum = lineSum + dataAsNum
						found = true
						break PosLoop
					}
				}
				if (strings.ContainsAny(string(currentLine[i]),charList)) {
					dataAsNum, _ := strconv.Atoi(data)
					//fmt.Println("Found Cur:",dataAsNum)
					lineSum = lineSum + dataAsNum
					found = true
					break PosLoop
				}
			}
			if (!found){
				//fmt.Println("Data:",data,"\nPL:",prevLine,"\nCL:",currentLine,"\nNL:",nextLine,"Start:",startPos,"End:",endPos)
			}
	}

	if (version == 2){
	}

	return lineSum
}

