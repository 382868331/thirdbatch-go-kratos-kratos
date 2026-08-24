package ewma

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKratos004SourceContract(t *testing.T) {
    source, err := os.ReadFile("node.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if n.errHandler != nil && n.errHandler(di.Err) {") {
        t.Fatalf("expected source contract is missing")
    }
}
