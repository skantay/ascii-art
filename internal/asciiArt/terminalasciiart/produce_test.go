package terminalasciiart

import (
	"os"
	"testing"
)

var testCases = []struct {
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
	{"test case #16", "", "standard.txt", "test16.txt"},
	{"test case #17", "\n", "standard.txt", "test17.txt"},
	{"test case #18", "hello\n", "standard.txt", "test18.txt"},
}

func TestConverter(t *testing.T) {
	for _, testCase := range testCases {
		want, _ := os.ReadFile("testCases/" + testCase.testFile)
		got := converter(testCase.input, testCase.banner)

		if string(want) != got {
			t.Fatalf("Test:\n%s\nGot:\n%sWant:\n%s", testCase.name, got, string(want))
		}
	}
}

func BenchmarkConverter(b *testing.B) {
	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			converter(tc.input, tc.banner)
		}
	}
}
