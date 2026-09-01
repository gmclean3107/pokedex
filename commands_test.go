package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCommandExit(t *testing.T) {
	if os.Getenv("TEST_EXIT") == "1" {
		commandExit(&config)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCommandExit")

	cmd.Env = append(os.Environ(), "TEST_EXIT=1")

	var output bytes.Buffer
	cmd.Stdout = &output

	err := cmd.Run()

	if err != nil {
		t.Fatalf("commandExit returned an error: %v", err)
	}

	expected := "Closing the Pokedex... Goodbye!\n"

	if output.String() != expected {
		t.Errorf("expected %q, got %q", expected, output.String())
	}
}

func TestHelpCommand(t *testing.T) {
	expected := []string{}

	for k := range CommandRegistry {
		expected = append(expected, k)
	}

	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	commandHelp(&config)

	w.Close()

	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()

	for _, command := range expected {
		if !strings.Contains(output, command) {
			t.Errorf("Actual is missing command(s)")
		}
	}
}
