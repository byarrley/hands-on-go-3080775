// challenges/testing/begin/main_test.go
package main

import "testing"

/*
Task:
- see challenge.md
- symbolCounter uses the buggy version of the function
Notes:
- If the "counter" interface still existed, this might've
	had a single test, run over a loop of counters using the
	same set of test cases
	(or maybe a map of types to specialized test cases)
Solution notes:
- I missed the "empty" case
- No need to set the identifier/designation/label fields in the counters (they're unused, so can be the zero value)
- The input name shown in the challenge.md's sample output doesn't match the actual inputs in the code (no underscores in the code)
- "#00" isn't a case for letterCount, seems like that's what's printed as the name for the zero value of string
- I'm leaving the "name" field for testCase, even though the instructor used the input itself

*/

type testCase struct {
	name string
	in   string
	want int
}

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	c := letterCounter{identifier: "letters"}
	tcs := []testCase{
		{name: "#00", in: "#00", want: 0},
		{name: "a2_32_^_&/)", in: "a2_32_^_&/)", want: 1},
		{name: "812_%6ab//", in: "812_%6ab//", want: 2},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			got := c.count(tc.in)
			if got != tc.want {
				t.Errorf("Want %d, got %d", tc.want, got)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	c := numberCounter{designation: "numbers"}
	tcs := []testCase{
		{name: "#00", in: "#00", want: 2},
		{name: "abc_1,?/", in: "abc_1,?/", want: 1},
		{name: "abc_12&8_^", in: "abc_12&8_^", want: 3},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			got := c.count(tc.in)
			if got != tc.want {
				t.Errorf("Want %d, got %d", tc.want, got)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	c := symbolCounter{label: "symbols"}
	tcs := []testCase{
		{name: "#00", in: "#00", want: 1},
		{name: "abc_1,?/", in: "abc_1,?/", want: 4},
		{name: "abc_12&8_^", in: "abc_12&8_^", want: 4},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			got := c.count(tc.in)
			if got != tc.want {
				t.Errorf("Want %d, got %d", tc.want, got)
			}
		})
	}
}
