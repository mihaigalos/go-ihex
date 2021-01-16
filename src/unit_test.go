package main

import (
	"testing"
)

func TestNumberOfBytesWorks_whenTypical(t *testing.T) {
	expected := 10

	actual := NumberOfBytes(":0A0000000EC015C014C013C012C011")

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}

func TestStartingAddressWorks_whenTypical(t *testing.T) {
	expected := 32256

	actual := StartingAddress(":107E000011E0A0E0B1E0E0E1F0E802C005900D92E1")

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}

func TestRecordWorks_whenTypical(t *testing.T) {
	expected := 171

	actual := Record(":107E00AB11E0A0E0B1E0E0E1F0E802C005900D92E1")

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}

func TestPayloadWorks_whenTypical(t *testing.T) {
	expected := [32]int{17, 224, 160, 224, 177, 224, 224, 225, 240, 232, 2, 192, 5, 144, 13, 146}

	actual_payload := Payload(":107E000011E0A0E0B1E0E0E1F0E802C005900D92E1")

	for i, e := range actual_payload {

		if e != int(expected[i]) {
			t.Errorf("No Match: %d != %d", e, expected[i])
		}

	}
}

func TestCRCWorks_whenTypical(t *testing.T) {
	expected := 225

	actual := CRC(":107E00AB11E0A0E0B1E0E0E1F0E802C005900D92E1")

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}
