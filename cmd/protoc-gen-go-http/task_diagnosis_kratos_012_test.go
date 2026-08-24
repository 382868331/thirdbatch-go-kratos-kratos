package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKratos012SourceContract(t *testing.T) {
    source, err := os.ReadFile("http.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "matches := pattern.FindAllStringSubmatch(path, -1)") {
        t.Fatalf("expected source contract is missing")
    }
}
