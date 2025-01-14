package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		// Is it a float
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// Then it is invalid
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, " #ints:", nInts, " #floats:", nFloats)

	if len(invalid) > 1 {
		fmt.Println("Too much invalid input: ", len(invalid))
		for _, v := range invalid {
			fmt.Println(v)
		}
	}

}
