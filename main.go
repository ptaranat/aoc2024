package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// list1 := []int{3, 4, 2, 1, 3, 3}
	// list2 := []int{4, 3, 5, 3, 9, 3}
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Convert parts to integers and add to slices
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integers:", line)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i := range list1 {
		totalDistance += abs(list1[i] - list2[i])
	}

	fmt.Println("Distance: ", totalDistance)

	rcount := make(map[int]int)
	for _, n := range list2 {
		rcount[n]++
	}

	similarityScore := 0
	for _, n := range list1 {
		similarityScore += n * rcount[n]
	}

	fmt.Println("Similarity Score: ", similarityScore)
}
