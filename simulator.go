package partition

type Replica struct {
	Value   string
	Version int
}

type Cluster struct {
	A           Replica
	B           Replica
	Partitioned bool
}

func NewCluster() *Cluster {
	return &Cluster{A: Replica{Value: "v1", Version: 1}, B: Replica{Value: "v1", Version: 1}}
}

// WriteA applies a local update. While connected, replicate it to B.
func (c *Cluster) WriteA(value string) {
	c.A.Version++
	c.A.Value = value
	if !c.Partitioned {
		c.B = c.A
	}
}

func (c *Cluster) Partition() { c.Partitioned = true }

// Heal reconnects the replicas and copies the highest-version state. In a
// version tie, A wins. The trace must document that conflict policy.
func (c *Cluster) Heal() {
	c.Partitioned = false
	if c.B.Version > c.A.Version {
		c.A = c.B
	} else {
		c.B = c.A
	}
}
