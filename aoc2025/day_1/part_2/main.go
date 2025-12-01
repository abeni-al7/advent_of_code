package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	password := ProcessFile("../day_1_input.txt")
	fmt.Println(password)
	fmt.Println()
}

func ProcessFile(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := ScanLineByLine(file)
	password := ProcessLines(lines)
	defer file.Close()
	return password
}

func ScanLineByLine(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ProcessLines(lines []string) int {
	curr := 50
	pass := 0
	for _, line := range lines {
		dir := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == "R" {
			pass += (int(math.Floor(float64(curr + amount) / float64(100)))) - (int(math.Floor(float64(curr) / float64(100))))
			curr += amount
		} else {
			pass += (int(math.Floor(float64(curr - 1) / float64(100)))) - (int(math.Floor(float64(curr - amount - 1) / float64(100))))
			curr -= amount
		}
		curr %= 100
		if curr < 0 {
			curr += 100
		}
	}
	return pass
}