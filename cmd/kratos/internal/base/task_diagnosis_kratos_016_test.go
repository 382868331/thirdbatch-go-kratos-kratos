package base

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKratos016SourceContract(t *testing.T) {
    source, err := os.ReadFile("mod.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if strings.Contains(str, path+\"@\") && i != -1 {") {
        t.Fatalf("expected source contract is missing")
    }
}
