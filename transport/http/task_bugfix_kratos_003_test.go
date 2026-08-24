package http

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos003SourceContract(t *testing.T) {
    source, err := os.ReadFile("stream.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if res != nil && res.Body != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
