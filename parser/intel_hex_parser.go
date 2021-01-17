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
		panic(".")
	}

	return int(value)

}

func StartingAddress(line string) int {
	slice := line[3:7]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
		panic(".")
	}

	return int(value)
}

func Record(line string) int {
	slice := line[7:9]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
		panic(".")
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
			panic(".")
		}
	}

	return result
}

func CRC(line string) uint8 {
	n := 2 * NumberOfBytes(line)
	slice := line[9+n : 9+n+2]

	value, err := strconv.ParseInt(slice, 16, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
		panic(".")
	}

	return uint8(int8(value))
}

func IsCRCValid(line string) bool {
	n := NumberOfBytes(line)
	a := StartingAddress(line)
	r := Record(line)
	p := Payload(line)
	expectedCRC := CRC(line)

	computedCRC := uint8(n + a + r)
	for _, e := range p {
		computedCRC = computedCRC + uint8(e)
	}

	return -int8(computedCRC) == int8(expectedCRC)
}

func IsFileValid(file []string) bool {
	for _, line := range file {

		actual := IsCRCValid(line)
		if actual == false {
			return false
		}

	}
	return true
}