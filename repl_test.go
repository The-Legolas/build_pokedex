package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v: expected word vs actual %v vs %v", c.input, actual, c.expected, expectedWord, word)
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()

	exitCmd, ok := commands["exit"]
	if !ok {
		t.Errorf("exit command not found in registry")
	}

	expectedDesc := "Exit the Pokedex"
	if exitCmd.description != expectedDesc {
		t.Errorf("expected description %v, got %v", expectedDesc, exitCmd.description)
	}

	helpCmd, ok := commands["help"]
	if !ok {
		t.Errorf("help command not found in registry")
	}

	expectedDesc = "Displays a help message"
	if helpCmd.description != expectedDesc {
		t.Errorf("expected description %v, got %v", expectedDesc, helpCmd.description)
	}
}
