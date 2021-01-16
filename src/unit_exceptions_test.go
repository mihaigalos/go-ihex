package main

import (
	"testing"
)

func TestNumberOfBytesWorks_whenException(t *testing.T) {
	defer func() { recover() }()
	NumberOfBytes(":XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	t.Errorf("should have panicked")
}

func TestStartingAddressWorks_whenException(t *testing.T) {
	defer func() { recover() }()
	StartingAddress(":XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	t.Errorf("should have panicked")
}

func TestRecordWorks_whenException(t *testing.T) {
	defer func() { recover() }()
	Record(":XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	t.Errorf("should have panicked")
}

func TestPayloadWorks_whenException(t *testing.T) {
	defer func() { recover() }()
	Payload(":10XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	t.Errorf("should have panicked")
}

func TestCRCWorks_whenException(t *testing.T) {
	defer func() { recover() }()
	CRC(":10XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	t.Errorf("should have panicked")
}
