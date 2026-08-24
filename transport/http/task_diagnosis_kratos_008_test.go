package http

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKratos008SourceContract(t *testing.T) {
    source, err := os.ReadFile("stream.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if s.conn == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
