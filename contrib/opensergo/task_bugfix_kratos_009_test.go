package opensergo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos009SourceContract(t *testing.T) {
    source, err := os.ReadFile("opensergo.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
