package consul

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos018SourceContract(t *testing.T) {
    source, err := os.ReadFile("registry.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if ss == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && ss == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
