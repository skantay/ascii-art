package conversion

import (
	"os"
	"testing"
)

func TestConvertText(t *testing.T) {
	testCases := []struct {
		input      string
		fileTarget string
	}{
		{"", "./testCases/target/target-1.txt"},
		{"hello", "testCases/target/target-2.txt"},
		{"HELLO", "testCases/target/target-3.txt"},
		{"HeLlo HuMaN", "testCases/target/target-4.txt"},
		{"1Hello 2There", "testCases/target/target-5.txt"},
		{"Hello\nThere", "testCases/target/target-6.txt"},
		{"Hello\n\nThere", "testCases/target/target-7.txt"},
		{"{Hello & There #}", "testCases/target/target-8.txt"},
		{"hello There 1 to 2!", "testCases/target/target-9.txt"},
		{"MaD3IrA&LiSboN", "testCases/target/target-10.txt"},
		{"1a\"#FdwHywR&/()=", "testCases/target/target-11.txt"},
		{"{|}~", "testCases/target/target-12.txt"},
		{`[\]^_ 'a`, "testCases/target/target-13.txt"},
		{"RGB", "testCases/target/target-14.txt"},
		{":;<=>?@", "testCases/target/target-15.txt"},
		{`\!" #$%&'"'"'()*+,-./`, "testCases/target/target-16.txt"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "testCases/target/target-17.txt"},
		{"abcdefghijklmnopqrstuvwxyz", "testCases/target/target-18.txt"},
	}

	for _, testCase := range testCases {
		got := ConvertText(testCase.input)
		buffer, err := os.ReadFile(testCase.fileTarget)
		if err != nil {
			t.Fatalf("error %s", err)
		}

		want := string(buffer)

		if got != want {
			t.Fatalf("\nGot:\n\n%s\n\nWant\n\n%s", got, want)
		}
	}
}
