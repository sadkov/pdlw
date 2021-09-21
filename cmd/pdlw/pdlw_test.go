package main

import (
	"testing"
)

const testCharList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZабцファ"

var testDict = map[int]string{
	1: "Sy2nMkPhcvvugTT1vn7zMnl3uODnJRcフフフmmmmmmm",
	2: "764918762349876",
	3: "Dp1A0zY5wrtYTBiOmBNol",
	4: "Zp-VNQcsN4DoU9KbMbHKVAqkz",
	5: "KyzjdH0NG7f6DL-vZuPIY5iVqu9Da",
	6: "a1a1a1a1a1a1a1a1a1a1a1a1a1b",
	7: "a1a1a1a1a1a1a1a1a1a1a1a1b1",
	8: "",
}

func TestLongestNone(t *testing.T) {
	name := "TestLongestNone"
	//expected := "ok"
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		//t.Logf("string to Map conversion error in %s", name)
		t.Fatalf("string to Map conversion error in %s", name)
		//t.Errorf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[2], runeMap)
	if err == nil {
		t.Errorf("%s: expected: error, got: %v", name, longest)
	}
}

func TestLongestLastAscii(t *testing.T) {
	name := "TestLongestLastAscii"
	expected := "KbMbHKVAqkz"
	dictTestEntry := 4
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		t.Fatalf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[dictTestEntry], runeMap)
	if err != nil {
		t.Errorf("%s: expected: %s, got error: %v", name, expected, err)
	} else if longest != expected {
		t.Errorf("%s: expected: %s, got longest: %v", name, expected, longest)
	}
}

func TestLongestLastUTF8(t *testing.T) {
	name := "TestLongestLastUTF8"
	expected := "uODnJRcフフフmmmmmmm"
	dictTestEntry := 1
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		t.Fatalf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[dictTestEntry], runeMap)
	if err != nil {
		t.Errorf("%s: expected: %s, got error: %v", name, expected, err)
	} else if longest != expected {
		t.Errorf("%s: expected: %s, got longest: %v", name, expected, longest)
	}
}

func TestLongestManyALastB(t *testing.T) {
	name := "TestLongestManyALastA"
	expected := "b"
	dictTestEntry := 6
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		t.Fatalf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[dictTestEntry], runeMap)
	if err != nil {
		t.Errorf("%s: expected: %s, got error: %v", name, expected, err)
	} else if longest != expected {
		t.Errorf("%s: expected: %s, got longest: %v", name, expected, longest)
	}
}

func TestLongestManyALastDigit(t *testing.T) {
	name := "TestLongestManyALastDigit"
	expected := "b"
	dictTestEntry := 6
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		t.Fatalf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[dictTestEntry], runeMap)
	if err != nil {
		t.Errorf("%s: expected: %s, got error: %v", name, expected, err)
	} else if longest != expected {
		t.Errorf("%s: expected: %s, got longest: %v", name, expected, longest)
	}
}

func TestLongestZeroValue(t *testing.T) {
	name := "TestLongestZeroValue"
	dictTestEntry := 8
	var runeMap = make(map[rune]bool)

	if runeMap = charToMap(testCharList); runeMap == nil {
		t.Fatalf("string to Map conversion error in %s", name)
	}

	longest, err := findLongest(testDict[dictTestEntry], runeMap)
	if err == nil {
		t.Errorf("%s: expected: error, got: %v", name, longest)
	}
}
