package wrr

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos005SourceContract(t *testing.T) {
    source, err := os.ReadFile("wrr.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
