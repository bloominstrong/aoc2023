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
		sum = sum + mySum(string(scanner.Bytes()))
		fmt.Println(sum)
	}
	//test := "eight9fhstbssrplmdlncmmqqnklb39ninejz"
	fmt.Println(sum)	
}

func mySum (s string ) int {

	first := strings.IndexAny(s,"1234567890")
	last  := strings.LastIndexAny(s,"1234567890")
	if (first == -1 || last == -1) {
		return 0
	}
	//fmt.Println("F:",first,"L:",last)
	ab := strings.Join([]string{string(s[first]),string(s[last])},"")
	AB, _ := strconv.Atoi(ab) 
	//fmt.Println("AB is ",AB)
	return AB
	

}
