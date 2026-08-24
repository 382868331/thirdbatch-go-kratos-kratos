package etcd

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos002SourceContract(t *testing.T) {
    source, err := os.ReadFile("registry.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if registerErr != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if registerErr == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
