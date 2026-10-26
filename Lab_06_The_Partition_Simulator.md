# Lab 6: The Partition Simulator

**CS 411: Computing Paradigms - Spring 2027**  
**Week 7: February 22-28, 2027**  
**Due:** February 28, 2027  
**Value:** 100 points

## Purpose

Simulate two replicas before, during, and after a network partition. Implement explicit CP and AP policies, then measure both the response behavior and key movement under consistent hashing.

## Learning objectives

- Apply CAP only when a partition exists
- Trace stale and rejected reads
- Explain consistent-hash key movement

## Before you begin

1. Review the Week 7 slide deck and assigned VoiceThread or reading.
2. Create your own repository from the Spring 2027 template linked in Blackboard. It includes the replica model, public tests, and TODOs in `policies.go` and `hashring.go`.
3. Run `go test ./...` before editing and save the initial output in `Lab6_Terminal.txt`. The policy and ring checks fail at the intended TODOs.

## Part A Partition trace (30 points)

Record the healthy baseline before separating replicas.

- In `Lab6_Trace.md` or PDF, begin with both replicas at `v1`, version 1. **Partition A and B, then write `v2` to A.** If you write before the partition, B also receives `v2` and the stale-read case disappears.
- Trace a read from B under CP and AP while partitioned. Record the returned value, accepted/rejected state, and possible-staleness flag for each policy.
- Heal the partition and trace B's value again. State the supplied version rule: the higher version wins; A wins a tie.

## Part B Policy implementation (40 points)

Make the trade-off observable.

- Complete `ReadCP` so it rejects reads when the two replicas cannot form the modeled quorum, even if a local copy exists.
- Complete `ReadAP` so it returns the selected local copy and marks it possibly stale during a partition. A stale flag expresses uncertainty, not proof that the value differs.
- Use the supplied `Heal` operation to reconcile state, then explain what safety and availability each policy chose during the partition.

## Part C Hash-ring change (30 points)

Add or remove one node.

- Complete `Owner` by hashing each node ID and key with the supplied FNV-1a function and choosing the first node clockwise, wrapping around when necessary. Node order in the input slice must not affect ownership.
- Add `north` to the ring containing `birch`, `cedar`, and `oak`. For the ten public-test keys, record owners before and after in the trace and calculate `100 × moved / total` percent.
- Explain why only keys in the clockwise interval from the new node's predecessor to the new node should move. Name at least one unchanged key and one moved key from your table.

## Verification

- `go test ./...` and `go test -race ./...` pass; save complete commands/output in `Lab6_Terminal.txt`.
- The trace includes healthy, partitioned, and healed states in that order.
- The ring report lists every public-test key, its before/after owner, the moved count, and the percentage.

## What to submit

Submit your own repository URL and:

1. Your own repository URL (created from the Spring 2027 template).
2. `simulator.go`, `policies.go`, and `hashring.go` in that repository.
3. `Lab6_Trace.pdf` or `Lab6_Trace.md`, including the trade-off explanation and ring table.
4. `Lab6_Terminal.txt`.

## Rubric

| Criterion | Points |
| --- | ---: |
| Partition trace | 25 |
| CP policy | 20 |
| AP policy | 20 |
| Consistent-hash implementation | 20 |
| Trade-off explanation | 15 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required trace file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.

