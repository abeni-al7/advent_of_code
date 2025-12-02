package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	password := ProcessFile("../day_1_input.txt")
	fmt.Println(password)
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
			curr = ((curr + amount) + 100) % 100
		} else {
			curr = ((curr - amount) + 100) % 100
		}
		if curr == 0 {
			pass++
		}
	}
	return pass
}