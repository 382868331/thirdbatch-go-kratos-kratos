package polaris

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos006SourceContract(t *testing.T) {
    source, err := os.ReadFile("config.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if event, ok := <-w.event; ok {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if event, ok := <=-w.event; ok {") {
        t.Fatalf("mutated source contract is still present")
    }
}
