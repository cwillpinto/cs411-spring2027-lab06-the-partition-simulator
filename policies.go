package partition

type ReadResult struct {
	Value         string
	Accepted      bool
	PossiblyStale bool
}

// ReadCP requires both replicas reachable, even when the local copy exists.
// TODO: reject a read during a partition; otherwise return the selected copy.
func (c *Cluster) ReadCP(replica string) ReadResult {
	return ReadResult{Value: c.local(replica).Value, Accepted: true}
}

// ReadAP always returns a local value. TODO: mark it possibly stale when the
// replicas are partitioned; the caller cannot verify the other copy.
func (c *Cluster) ReadAP(replica string) ReadResult {
	return ReadResult{Value: c.local(replica).Value, Accepted: true}
}

func (c *Cluster) local(replica string) Replica {
	if replica == "B" {
		return c.B
	}
	return c.A
}
