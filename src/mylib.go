package main

import (
	"fmt"
	"strconv"
)

func NumberOfBytes(line string) int {
	slice := line[1:3]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
	}

	return int(value)

}

func StartingAddress(line string) int {
	slice := line[3:7]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
	}

	return int(value)
}

func Record(line string) int {
	slice := line[7:9]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
	}

	return int(value)
}
