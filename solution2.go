package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X, Y int
}

type PosCorrected struct {
	X, Y, Aim int
}

func solutionDayTwoPartOne() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := Pos{}

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		instrcution := strings.Split(line, " ")
		mag, _ := strconv.Atoi(instrcution[1])

		switch instrcution[0] {
		case "forward":
			pos.X = pos.X + mag
		case "up":
			pos.Y = pos.Y - mag
		case "down":
			pos.Y = pos.Y + mag
		}
	}

	fmt.Println("X ", pos.X, "Y ", pos.Y, "mul ", pos.X*pos.Y)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solutionDayTwoPartTwo() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := PosCorrected{}

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		instrcution := strings.Split(line, " ")
		mag, _ := strconv.Atoi(instrcution[1])

		switch instrcution[0] {
		case "forward":
			pos.X = pos.X + mag
			pos.Y = pos.Y + (pos.Aim * mag)
		case "up":
			pos.Aim = pos.Aim - mag
		case "down":
			pos.Aim = pos.Aim + mag
		}
	}

	fmt.Println("X ", pos.X, "Y ", pos.Y, "mul ", pos.X*pos.Y)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
