package terminalasciiart_test

import (
	"testing"

	"github.com/skantay/ascii-art/internal/asciiArt/terminalasciiart"
)

// Test ValidateFile function

// Positive
func TestValidateFilePositive(t *testing.T) {
	ValidTestCases := []struct {
		testName string
		banner   string
		err      error
	}{
		{
			testName: "positive standard.txt",
			banner:   "standard.txt",
			err:      nil,
		},
		{
			testName: "positive shadow.txt",
			banner:   "shadow.txt",
			err:      nil,
		},
		{
			testName: "positive thinkertoy.txt",
			banner:   "thinkertoy.txt",
			err:      nil,
		},
	}

	for _, testCase := range ValidTestCases {
		err := terminalasciiart.ValidateBanner(testCase.banner)

		if err != testCase.err {
			t.Fatalf("Test name: %s\nGot:%s\nWant:%s", testCase.testName, err, testCase.err)
		}
	}
}

// Negative
func TestValidateFileNegative(t *testing.T) {
	InvalidTestCases := []struct {
		testName string
		banner   string
	}{
		{
			testName: "negative thinkertoy.txt",
			banner:   "thintoy.txt",
		},
		{
			testName: "negative shadow.txt",
			banner:   "invalid/shadow.txt",
		},
		{
			testName: "negative non banner",
			banner:   "",
		},
	}

	for _, testCase := range InvalidTestCases {
		err := terminalasciiart.ValidateBanner(testCase.banner)

		if err == nil {
			t.Fatalf("Test name: %s\nGot:%s\nExpected: error", testCase.testName, err)
		}
	}
}
