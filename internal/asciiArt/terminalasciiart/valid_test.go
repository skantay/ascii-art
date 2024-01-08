package terminalasciiart

import (
	"testing"
)

// Test ValidateFile function

// Positive
func TestValidateBannerPositive(t *testing.T) {
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
		err := validateBanner(testCase.banner)

		if err != testCase.err {
			t.Fatalf("Test name: %s\nGot:%s\nWant:%s", testCase.testName, err, testCase.err)
		}
	}
}

// Negative
func TestValidateBannerNegative(t *testing.T) {
	InvalidTestCases := []struct {
		testName string
		banner   string
	}{
		{
			testName: "banner not found",
			banner:   "sha.txt",
		},
	}

	for _, testCase := range InvalidTestCases {
		err := validateBanner(testCase.banner)
		if err == nil {
			t.Fatalf("\nTest name:%s\nExpected error:%s\n", testCase.testName, err)
		}
	}
}

// Test ValidateInput function

// Positive
func TestValidateInputPositive(t *testing.T) {
	ValidTestCases := []struct {
		text string
		err  error
	}{
		{
			text: "qwertyuiopasdfghjklzxcvbnm",
			err:  nil,
		},
		{
			text: "234567890-=",
			err:  nil,
		},
		{
			text: "!@#$%^&*()_{:>}><?L",
			err:  nil,
		},
	}

	for i, testCase := range ValidTestCases {
		err := validateInput(testCase.text)

		if err != testCase.err {
			t.Fatalf("Test index: %d\nGot:%s\nWant:%s", i, err, testCase.err)
		}
	}
}

// Negative
func TestValidateInputNegative(t *testing.T) {
	InvalidTestCases := []struct {
		text string
	}{
		{
			text: "抄本《東明聞見錄》頁面順序有錯",
		},
		{
			text: "Я встаю очень рано в шесть утра по звонку будильника",
		},
		{
			text: "ومة من قبل اغلب برامج التصميم مثل الفوتوش",
		},
	}

	for i, testCase := range InvalidTestCases {
		err := validateInput(testCase.text)

		if err == nil {
			t.Fatalf("Test index: %d\nGot:%s\nExepcted: error", i, err)
		}
	}
}
