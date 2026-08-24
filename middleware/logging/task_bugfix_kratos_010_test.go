package logging

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos010SourceContract(t *testing.T) {
    source, err := os.ReadFile("logging.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
