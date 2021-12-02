package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solutionDayOnePartOne() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	oldValue := -1
	count := 0

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		val := scanner.Text()
		intval, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		if oldValue != -1 && intval > oldValue {
			fmt.Print(count, intval, "\n")
			count = count + 1
		}

		oldValue = intval
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(count)
}

func solutionDayOnePartTwo() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sliceOfInts := []int{}
	for scanner.Scan() {
		val := scanner.Text()
		intval, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		sliceOfInts = append(sliceOfInts, intval)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(sliceOfInts)

	aggSlice := []int{}
	for i := 0; i <= len(sliceOfInts)-3; i++ {
		aggSlice = append(aggSlice, sliceOfInts[i]+sliceOfInts[i+1]+sliceOfInts[i+2])
	}

	oldValue := -1
	count := 0

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for k := range aggSlice {
		intval := aggSlice[k]

		if oldValue != -1 && intval > oldValue {
			fmt.Print(count, intval, "\n")
			count = count + 1
		}

		oldValue = aggSlice[k]
	}

	fmt.Print(count)
}
