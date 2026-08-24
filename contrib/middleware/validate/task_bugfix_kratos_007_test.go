package validate

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos007SourceContract(t *testing.T) {
    source, err := os.ReadFile("validate.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := v.Validate(); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err := v.Validate(); err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
