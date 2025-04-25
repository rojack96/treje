# Treje

A lightweight, idiomatic and extensible collection of data structures in Go.

> Treje aims to provide clean implementations of common data structures not available in the Go standard library — starting with Set & MapSet.

## Features (Current)

✅ Set implementation   
✅ MapSet implementation  
✅ Operations:
- Manipulation: `Add`, `Remove`, `Discard`, `Pop`
- Set operations: `Union`, `Intersection`, `Difference`, `SymmetricDifference`
- `IsSubsetOf`, `Equals`  

✅ Utilities:

- `Has()`
- `IsEmpty()`
- `Clear()`
- `Min()` & `Max()`
- `Sum()` (numbers) or `Concat(separator)` (string) 
- `Sort()` & `ReverseSort()`
- `Copy()`
- `ToSlice()`

## Installation

```bash
go get github.com/rojack96/treje
```

## Example

```go
A := treje.NewSet().Int8(1, 2, 3)
A.Add(4)
B := treje.NewSet().Int8(3, 4, 5)

diff := A.Difference(B) // [1 2]

fmt.Println("Difference:", diff)
```

## Planned Additions

- [x] Set
- [x] MapSet (Set backed by map for performance)
- [ ] Stack
- [ ] Queue
- [ ] Deque
- [ ] Linked List
- [ ] Tree structures (BST, AVL, etc.)
- [ ] Graph
- [ ] Priority Queue / Heap

## Design Goals

- Idiomatic Go
- Datatype first
- No external dependencies
- Generic-friendly (Go 1.18+ ready)
- Focus on clarity and correctness
- Simple API surface

## License

MIT License