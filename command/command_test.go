package command

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

// Test expected to run correct arguments and flags.
func TestParseCorrectFlagsAndArgs(t *testing.T) {
	var tests = []struct {
		action string
		body   string
	}{
		{"-la", ""},
		{"-ld", ""},
		{"-lc", ""},
		{"-a", "New task description."},
		{"-c", "123"},
		{"-d", "123"},
		{"-r", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.body}
			_, err := parseFlagsAndArgs()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to not throw errors: %v", tt)
			}
		})
	}
}

// Test expected to catch incorrect arguments and flags.
func TestParseIncorrectFlagsAndArgs(t *testing.T) {
	var tests = []struct {
		action string
		body   string
	}{
		{"-a", ""},
		{"-c", "abc123"},
		{"-d", ""},
		{"-s", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.body}
			_, err := parseFlagsAndArgs()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to fail because of errors: %v", tt)
			}
		})
	}
}

// Test expected to catch extra arguments.
func TestExtraArgs(t *testing.T) {
	var tests = []struct {
		action   string
		body     string
		extraArg string
	}{
		{"-a", "-la", "wefwe"},
		{"-c", "123", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd with extra argument %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.body, tt.extraArg}
			_, err := parseFlagsAndArgs()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to fail for extra arguments: %v", tt)
			}
		})
	}
}
