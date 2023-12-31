package main

	import (
		"fmt"
		"strings"
		"strconv"
		"os"
		"bufio"
	)

func main() {
	//from input
	file, _ := os.Open("../input/1")
	scanner := bufio.NewScanner(file)
	sum := 0
	 
	for scanner.Scan() {
		//fmt.Println(string(scanner.Bytes()))
		sum = sum + mySum(string(scanner.Bytes()),1)
		//fmt.Println(sum)
	}
	file.Close()
	fmt.Println("Version 1 is:",sum)


	file, _ = os.Open("../input/1")
	scanner2 := bufio.NewScanner(file)
	sum = 0
	for scanner2.Scan() {
		//fmt.Println(string(scanner2.Bytes()))
		sum = sum + mySum(string(scanner2.Bytes()),2)
		//fmt.Println(sum)
	}
	fmt.Println("Version 2 is:",sum)
}

func mySum (s string ,version int) int {

	first := strings.IndexAny(s,"1234567890")
	last  := strings.LastIndexAny(s,"1234567890")
	
	if (first == -1 || last == -1) {
		return 0
	}
	a := string(s[first])
	b := string(s[last])
	if (version == 2 ) {
		firstv2, firstv2Val, lastv2, lastv2Val := findPrintedNumber(s)
		if (firstv2!=-1 && firstv2<first) {
			a = strconv.Itoa(firstv2Val)
		}
		if (lastv2!=-1 && lastv2>last) {
			b = strconv.Itoa(lastv2Val)
		}
		//fmt.Println("F:",first,"L:",last,"F2:",firstv2,"L2:",lastv2,"String:",s,"A:",a,"B:",b)
	}
	ab := strings.Join([]string{a,b},"")
	AB, _ := strconv.Atoi(ab) 
	//fmt.Println("AB is ",AB)
	return AB
}

func findPrintedNumber (s string) (int, int, int, int) {
	numbers := []string{"one","two","three","four","five","six","seven","eight","nine"}
	numbersVal := []int{1,2,3,4,5,6,7,8,9}
	firstPos, lastPos := -1, -1
	firstVal, lastVal := -1, -1

	for i, element := range numbers {
		if (strings.Contains(s,element)) {
			var elementFirstPos int = strings.Index(s,element)
			if (firstPos == -1 || elementFirstPos < firstPos) {
				firstPos = elementFirstPos
				firstVal = numbersVal[i]
			}

			var elementLastPos int = strings.LastIndex(s,element)
			if (lastPos == -1 || elementLastPos > lastPos) {
				lastPos = elementLastPos
				lastVal = numbersVal[i]
			}
		}
	}
	return firstPos, firstVal, lastPos, lastVal
}
