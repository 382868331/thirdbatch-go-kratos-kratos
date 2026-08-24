package form

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos017SourceContract(t *testing.T) {
    source, err := os.ReadFile("proto_decode.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
