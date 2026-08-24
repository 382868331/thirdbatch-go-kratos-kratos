package subset

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKratos013SourceContract(t *testing.T) {
    source, err := os.ReadFile("subset.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if num <= 0 || len(inss) <= num {") {
        t.Fatalf("expected source contract is missing")
    }
}
