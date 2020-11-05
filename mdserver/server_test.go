package mdserver

import (
	"fmt"
	"testing"
)

func displayArr(arr *[]string) string {
	s := "{"
	for _, x := range *arr {
		s += fmt.Sprintf(`"%v", `, x)
	}
	s += "}"

	return s
}

func printArrs(want, got *[]string, t *testing.T) {
	t.Errorf("Error: want `%v`, got `%v`", displayArr(want), displayArr(got))
}

func compArr(want, got []string, t *testing.T) {
	for i := range want {
		if i >= len(got) {
			printArrs(&want, &got, t)
		}

		if want[i] != got[i] {
			printArrs(&want, &got, t)
		}
	}

	if len(got) != len(want) {
		printArrs(&want, &got, t)
	}
}

func TestUndaemonArgs(t *testing.T) {
	input := []string{"gomd", "--quiet", "--daemon", "--no-open"}
	want := []string{"gomd", "--quiet", "--no-open"}
	got := undaemonArgs(&input)
	compArr(want, got, t)

	input = []string{"gomd", "-daemon", "--no-open", "--quiet"}
	want = []string{"gomd", "--no-open", "--quiet"}
	got = undaemonArgs(&input)
	compArr(want, got, t)

	input = []string{"gomd", "-daemon", "--daemon"}
	want = []string{"gomd"}
	got = undaemonArgs(&input)
	compArr(want, got, t)

}
