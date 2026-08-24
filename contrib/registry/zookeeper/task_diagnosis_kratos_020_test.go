package zookeeper

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKratos020SourceContract(t *testing.T) {
    source, err := os.ReadFile("register.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "case err = <-ch:") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "case err = <=-ch:") {
        t.Fatalf("mutated source contract is still present")
    }
}
