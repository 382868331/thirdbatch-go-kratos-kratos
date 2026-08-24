package log

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos015SourceContract(t *testing.T) {
    source, err := os.ReadFile("filter.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(cfg.keys) == 0 && cfg.filter == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
