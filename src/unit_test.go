package main

import "testing"

func TestNumberOfBytesWorks_whenTypical(t *testing.T) {
	expected := 10

	actual := NumberOfBytes(":0A0000000EC015C014C013C012C011")

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}
