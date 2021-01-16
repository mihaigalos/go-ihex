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

func TestCRCWorks_whenTypicalPositive(t *testing.T) {
	expected_in_twos_complement := uint8(0xE1)

	actual := CRC(":107E000011E0A0E0B1E0E0E1F0E802C005900D92E1")

	if uint8(actual) != uint8(expected_in_twos_complement) {
		t.Errorf("No Match: %d != %d", actual, expected_in_twos_complement)
	}
}
func TestCRCWorks_whenTypicalNegative(t *testing.T) {
	expected_in_twos_complement := uint8(0x1E)

	actual := CRC(":0300300002337A1E")

	if uint8(actual) != expected_in_twos_complement {
		t.Errorf("No Match: %d != %d", actual, expected_in_twos_complement)
	}
}

func TestIsCRCValidWorks_whenTypicalPositive(t *testing.T) {
	expected := true

	actual := IsCRCValid(":0300300002337A1E")

	if actual != expected {
		t.Errorf("No Match: %t != %t", actual, expected)
	}
}

func TestIsCRCValidWorks_whenTypicalNegative(t *testing.T) {
	expected := true

	actual := IsCRCValid(":1000800005900D92A030B107D9F711E0A0E0B1E0E2")

	if actual != expected {
		t.Errorf("No Match: %t != %t", actual, expected)
	}
}

func TestIsCRCOnWholeFileValid_whenTypical(t *testing.T) {
	expected := true

	file := []string{
		":100000000C9434000C944F000C944F000C944F004F",
		":100010000C944F000C944F000C944F000C944F0024",
		":100020000C944F000C944F000C944F000C944F0014",
		":100030000C944F000C944F000C944F000C944F0004",
		":100040000C944F000C944F000C944F000C944F00F4",
		":100050000C944F000C944F000C944F000C944F00E4",
		":100060000C944F000C944F0011241FBECFEFD4E02E",
		":10007000DEBFCDBF11E0A0E0B1E0E8EFF0E002C0EC",
		":1000800005900D92A030B107D9F711E0A0E0B1E0E2",
		":1000900001C01D92A030B107E1F70C9467000C94E9",
		":1000A00000008FEF84B987B98EEF8AB9089501C037",
		":1000B0000197009759F020E00000000000000000C8",
		":1000C000000000002F5F2A3599F3F6CF08958FEFD7",
		":1000D00084B987B98EEF8AB98FEF88B985B98BB9A2",
		":1000E00084EF91E00E94570018B815B81BB884EF50",
		":0800F00091E00E945700F0CFDF",
		":00000001FF"}

	actual := IsFileValid(file)

	if actual != expected {
		t.Errorf("No Match: %t != %t", actual, expected)
	}

}
