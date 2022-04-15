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
		command     string
		action      string
		body        string
		commandType CommandType
	}{
		{"command", "-la", "", ListAll},
		{"command", "-ld", "", ListDoing},
		{"command", "-lc", "", ListCompleted},
		{"command", "-a", "New task description.", AddTask},
		{"command", "-c", "123", CompleteTask},
		{"command", "-d", "123", DeleteTask},
		{"command", "-r", "", ResetTasks},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.command, tt.action, tt.body}
			command, err := ParseFromShell()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err != nil {
				t.Errorf("Expected command to not throw errors: %v", tt)
			}
			if command.Type != tt.commandType {
				t.Errorf("Expected command of type %v, got %v", tt.commandType, command.Type)
			}
		})
	}
}

// Test expected to catch extra arguments.
func TestExtraArgs(t *testing.T) {
	var tests = []struct {
		command  string
		action   string
		body     string
		extraArg string
	}{
		{"command", "-a", "la", "wefwe"},
		{"command", "-c", "123", ""},
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

// Test expected to catch incorrect arguments and flags.
func TestParseAndValidateIncorrectFlagsAndArgs(t *testing.T) {
	var tests = []struct {
		command string
		action  string
		body    string
	}{
		{"command", "-a", ""},
		{"command", "-c", "abc123"},
		{"command", "-d", ""},
		{"command", "-s", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.body}
			_, err := ParseFromShell()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to fail because of errors: %v", tt)
			}
		})
	}
}
