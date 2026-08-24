package discovery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos011SourceContract(t *testing.T) {
    source, err := os.ReadFile("resolver.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
