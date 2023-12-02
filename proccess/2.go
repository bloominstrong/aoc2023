package main

	import (
		"fmt"
		"strings"
		"strconv"
		"os"
		"bufio"
		"slices"
	)

func main() {
	//from input
	file, _ := os.Open("../input/2")
	scanner := bufio.NewScanner(file)
	sum := 0
	 
	for scanner.Scan() {
		//fmt.Println(string(scanner.Bytes()))
		sum = sum + mySum(string(scanner.Bytes()),1)
		//fmt.Println(sum)
	}
	file.Close()
	fmt.Println("Version 1 is:",sum)

	
	file, _ = os.Open("../input/2")
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
	maxGreen, maxBlue, maxRed := 13, 14, 12
	gameValid := false 
	var greens, reds, blues []int
	gameString, gameData, isStringValid := strings.Cut(s,":")
	var cubeData []string

	if (!isStringValid) {
		//input invalid
		return 0
	}
	gameIndex , _ := strconv.Atoi(strings.TrimLeft(gameString, "Game "))
	//fmt.Println("Game:",gameIndex)

	setString := strings.Split(gameData,";")
	for _, setData := range setString {
		cubeData = strings.Split(setData,",")
		//fmt.Println("Cubes:", cubeData)
		for _, cubes := range cubeData {
			cubeString := strings.Split(cubes," ")
			number , _ := strconv.Atoi(cubeString[1])
			//fmt.Println("Color, num:",cubeString[2],number)
			switch cubeString[2] {
			case "green": 
					greens = append(greens, number)
			case "red"  : 
					reds = append(reds, number)
			case "blue" : 
					blues = append(blues, number)
			default : fmt.Println("Error in cube parsing")
			}
		}
		//fmt.Println(greens,blues,reds)		
	}
	if (slices.Max(greens)>maxGreen||slices.Max(blues)>maxBlue||slices.Max(reds)>maxRed) {
		//fmt.Println("Invalid Game")
		gameValid = false
	} else {
		//fmt.Println("Valid Game")
		gameValid = true
	}
	if (version == 2) {
		cubePower := slices.Max(greens)*slices.Max(reds)*slices.Max(blues)
		return cubePower
	}
	if (gameValid) {
		return gameIndex
	} else {
		return 0
	}
}

