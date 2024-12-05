package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open input file
	file, err := os.Open("input.txt")

	// raise err if an err is returned from open and exit immediately
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
		os.Exit(1)
	}

	// defer file close to end of function so we can read it
	defer file.Close()

	var left []int
	var right []int

	// bufio puts data into a buffer instead of reading from disk every time
	scanner := bufio.NewScanner(file)
	// .Scan() will read in the passed in file and move it token by token in our case
	// default is line by line, so we read each number of lines ...
	for scanner.Scan() {
		line := scanner.Text()

		// take the string and split by whitespace .Fields does it by whitespace by default
		// we assume input is correct
		parts := strings.Fields(line)

		// strconv = package for converting strings which we're reading in from the file
		// and converting them to other types
		// in our case we use Atoi which is converting ascii to integer,
		// Atoii = "ascii to integer"
		// we assume valid input, so no errors to be raised so we throw away err
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		// append numbers
		left = append(left, num1)
		right = append(right, num2)
	}

	// we need this because scanner.Scan returns boolean of true/false if not successful is
	// false but no error is thrown so scanner.Err() allows us to see if there was an err
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Distance: %d\n", calcTotalDistance(left, right))
	fmt.Printf("Similarity Score: %d\n", calcSimilarityScore(left, right))
}

// runs in O(NLogN) for the sort, and then N for the calculation of total distance
// but worst case  is NlogN for sort, in place so O(1) for space
func calcTotalDistance(left, right []int) int {
	// sort both slices with builtin go package
	sort.Ints(left)
	sort.Ints(right)

	// calculate total distance which is abs(x1-x2) all totaled up for each line
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		// math.Abs takes in a float64 so we need to convert, then convert to int
		// concerns for loss of info?
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}

// runs in N time and space. N for the map that stores the ints, well N unique i guess
// and then time is M for running through all of left numbers
func calcSimilarityScore(left, right []int) int {
	// create map of frequency of numbers for the list on the right
	// so we know how many of left item exists in right for left*countOfRight
	rightFreqs := make(map[int]int)
	for _, x := range right {
		// ++ = shorthand syntax for incrementing count
		// equivalent to rightFreqs[x] = rightFreqs[x] + 1
		rightFreqs[x]++
	}

	// calc the similarity score
	var similarityScore int
	for _, x := range left {
		// check that it exists in the map, if it does update the similarity score
		if count, exists := rightFreqs[x]; exists {
			similarityScore += x * count
		}
	}

	return similarityScore
}
