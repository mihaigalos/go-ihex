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

func Payload(line string) [32]int {
	var result [32]int
	n := 2 * NumberOfBytes(line)
	slice := line[9 : 9+n]

	for i := 0; i < len(slice); i += 2 {
		hex := slice[i : i+2]
		value, err := strconv.ParseInt(hex, 16, 64)
		result[i/2] = int(value)

		if err != nil {
			fmt.Printf("Conversion failed: %s\n", err)
		}
	}

	return result
}

func CRC(line string) int {
	n := 2 * NumberOfBytes(line)
	slice := line[9+n : 9+n+2]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
	}

	return int(value)
}
