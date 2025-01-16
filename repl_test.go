package main

import "testing"

type CleanInputTestArgs struct {
	input string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []CleanInputTestArgs {
		{
			input: "  hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "CHarm   buLB ser",
			expected: []string{"charm", "bulb", "ser"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("amount of words mismatch")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word mismatch in index %d: %s != %s", i, word, expectedWord)
			}
		}
	}
}