package collector

import (
    "context"
    "testing"
)

func TestSystemCollector_Collect(t *testing.T) {
    c := NewSystemCollector()
    metrics, err := c.Collect(context.Background())
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if metrics.Timestamp.IsZero() {
        t.Error("expected timestamp to be set")
    }
    if metrics.CPU.Cores == 0 {
        t.Error("expected CPU cores > 0")
    }
}