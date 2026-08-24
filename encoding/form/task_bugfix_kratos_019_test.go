package form

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos019SourceContract(t *testing.T) {
    source, err := os.ReadFile("proto_encode.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for i := 0; i < len(s); i++ { // proto identifiers are always ASCIIS") {
        t.Fatalf("expected source contract is missing")
    }
}
