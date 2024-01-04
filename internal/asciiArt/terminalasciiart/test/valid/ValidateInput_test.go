package terminalasciiart_test

import (
	"testing"

	"github.com/skantay/ascii-art/internal/asciiArt/terminalasciiart"
)

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
		err := terminalasciiart.ValidateInput(testCase.text)

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
		err := terminalasciiart.ValidateInput(testCase.text)

		if err == nil {
			t.Fatalf("Test index: %d\nGot:%s\nExepcted: error", i, err)
		}
	}
}
