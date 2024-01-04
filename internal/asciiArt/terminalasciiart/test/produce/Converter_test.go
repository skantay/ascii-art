package terminalasciiart_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/skantay/ascii-art/internal/asciiArt/terminalasciiart"
)

func TestConverter(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		banner   string
		testFile string
	}{
		{"test case #1", "hello", "standard.txt", "test1.txt"},
		{"test case #2", "HELLO", "standard.txt", "test2.txt"},
		{"test case #3", "HeLlo HuMaN", "standard.txt", "test3.txt"},
		{"test case #4", "1Hello 2There", "standard.txt", "test4.txt"},
		{"test case #6", "{Hello & There #}", "standard.txt", "test6.txt"},
		{"test case #7", "hello There 1 to 2!", "standard.txt", "test7.txt"},
		{"test case #8", "MaD3IrA&LiSboN", "standard.txt", "test8.txt"},
		{"test case #9", `1a"#FdwHywR&/()=`, "standard.txt", "test9.txt"},
		{"test case #10", "{|}~", "standard.txt", "test10.txt"},
		{"test case #11", `[\]^_ 'a`, "standard.txt", "test11.txt"},
		{"test case #12", "RGB", "standard.txt", "test12.txt"},
		{"test case #13", ":;<=>?@", "standard.txt", "test13.txt"},
		{"test case #14", `\!" #$%&'()*+,-./`, "standard.txt", "test14.txt"},
		{"test case #15", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "standard.txt", "test15.txt"},
		{"test case #16", "abcdefghijklmnopqrstuvwxyz", "standard.txt", "test16.txt"},
	}

	for _, testCase := range testCases {
		want, _ := os.ReadFile("testCases/" + testCase.testFile)

		got := terminalasciiart.Converter(testCase.input, testCase.banner)

		if reflect.DeepEqual(got, string(want)) {
			t.Fatalf("\n%s\nGot:\n%s\nWant:\n%s\n", testCase.name, got, want)
		}
	}
}
