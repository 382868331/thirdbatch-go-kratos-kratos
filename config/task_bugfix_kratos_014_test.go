package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos014SourceContract(t *testing.T) {
    source, err := os.ReadFile("options.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if i == len(keys)-1 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if i != len(keys)-1 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
