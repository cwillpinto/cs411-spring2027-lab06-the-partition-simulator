package partition

import "testing"

func TestPolicyTrace(t *testing.T) {
	c := NewCluster()
	if c.ReadCP("B").Value != "v1" {
		t.Fatal("healthy baseline")
	}
	c.Partition()
	c.WriteA("v2")
	if got := c.ReadCP("B"); got.Accepted {
		t.Fatalf("CP accepted partitioned read: %+v", got)
	}
	if got := c.ReadAP("B"); !got.Accepted || got.Value != "v1" || !got.PossiblyStale {
		t.Fatalf("AP stale read: %+v", got)
	}
	c.Heal()
	if got := c.ReadCP("B"); !got.Accepted || got.Value != "v2" || got.PossiblyStale {
		t.Fatalf("healed read: %+v", got)
	}
}

func TestRingMovement(t *testing.T) {
	keys := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota", "kappa"}
	before := []string{"birch", "cedar", "oak"}
	after := []string{"birch", "cedar", "oak", "north"}
	moved, total := Movement(keys, before, after)
	if total != len(keys) || moved <= 0 || moved >= total {
		t.Fatalf("movement = %d/%d", moved, total)
	}
	for _, key := range keys {
		old, next := Owner(key, before), Owner(key, after)
		if old != next && next != "north" {
			t.Fatalf("key %q moved from %q to existing node %q", key, old, next)
		}
	}
}

func TestRingIndependentOfInputOrder(t *testing.T) {
	a, b := []string{"node-a", "node-b", "node-c"}, []string{"node-c", "node-a", "node-b"}
	for _, key := range []string{"alpha", "beta", "gamma"} {
		if Owner(key, a) != Owner(key, b) {
			t.Fatalf("owner depends on input ordering for %q", key)
		}
	}
}
