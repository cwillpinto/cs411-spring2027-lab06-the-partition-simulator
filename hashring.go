package partition

import "hash/fnv"

func Hash(value string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(value))
	return h.Sum32()
}

// Owner chooses the first node clockwise on the hash ring, wrapping around.
// TODO: implement this selection using Hash for the key and each node ID.
func Owner(key string, nodes []string) string {
	_ = key
	if len(nodes) == 0 {
		return ""
	}
	return nodes[0]
}

func Movement(keys, before, after []string) (moved int, total int) {
	for _, key := range keys {
		if Owner(key, before) != Owner(key, after) {
			moved++
		}
	}
	return moved, len(keys)
}
