# go-algorithm
Some algorithm functions found in the c++ STL developed with generics (golang 1.18).

# Install

```sh
go get github.com/AlexandreChamard/go-algorithm
```

# Examples

## Sorting

```golang
import (
    . "github.com/AlexandreChamard/go-builtin"
)

sliceInt := []int{42, 12, 4, -5, 0, 69, 20}
Sort(sliceInt, Less[int])
// sliceInt = []int{-5, 0, 4, 12, 20, 42, 69}
```

## Unique

```golang
import (
    . "github.com/AlexandreChamard/go-builtin"
)

type t struct {
    i, j int
}

eqf := func(a, b t) bool { return a.j == b.j }

slice := []t{{0, 42}, {1, 12}, {2, 4}, {3, -5}, {4, 0}, {5, 69}, {6, 4}, {7, 20}}

Unique(slice) // Returns true

UniqueF(slice, eqf) // Returns false
```
