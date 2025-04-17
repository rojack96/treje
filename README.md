# Treje

A lightweight, idiomatic and extensible collection of data structures in Go.

> Treje aims to provide clean implementations of common data structures not available in the Go standard library — starting with Set.

## ✨ Features (Current)

✅ Set implementation (e.g. `Int8Set`, `Uint32Set` `StringSet`)  
✅ Operations:
- Manipulation: `Add`, `Remove`, `Discard`, `Pop`, `IsEmpty`
- Set operations: `Union`, `Intersection`, `Difference`, `SymmetricDifference`
- `IsSubsetOf`, `Equals`  

✅ Utilities:

- `Has()`
- `IsEmpty()`
- `Clear()`
- `Min()`
- `Max()`
- `Max()`
- `Sum()` (for numbers) or `Concat(separator string)` for string 
- `Sort()`
- `ReverseSort()`
- `Copy()`
- `ToSlice()`

## 📦 Installation

```bash
go get github.com/rojack96/treje
```

## 🧪 Example

```go
A := set.NewInt8Set(1, 2, 3)
A.Add(4)
B := set.NewInt8Set(3, 4, 5)

diff := A.Difference(B) // [1 2]

fmt.Println("Difference:", diff)
```

## 📚 Planned Additions

- [x] Set
- [ ] MapSet (Set backed by map for performance)
- [ ] Stack
- [ ] Queue
- [ ] Deque
- [ ] Linked List
- [ ] Tree structures (BST, AVL, etc.)
- [ ] Graph
- [ ] Priority Queue / Heap

## 🔧 Design Goals

- Idiomatic Go
- No external dependencies
- Generic-friendly (Go 1.18+ ready)
- Focus on clarity and correctness
- Simple API surface

## 📄 License

MIT License