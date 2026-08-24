package p2c

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos001SourceContract(t *testing.T) {
    source, err := os.ReadFile("p2c.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return nodes[0], done, nil") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return nodes[1], done, nil") {
        t.Fatalf("mutated source contract is still present")
    }
}
