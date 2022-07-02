package main

import (
	"fmt"

	"github.com/b5710546232/sparseset/sparseset"
)

func main() {

	ss := sparseset.NewSparseSet[int](10, 10)
	ss.Contains(1)
	ss.Put(1, 99)
	ss.Contains(1)

	ss.Put(2, 10)

	ss.Put(10, 9)

	ss.Put(3, 7)
	ss.Put(3, 7)
	// ss.PrintSpase()
	// ss.PrintDense()
	ss.Remove(2)
	// ss.PrintSpase()
	// ss.PrintDense()

	// fmt.Println("ss.Contains(3)", ss.Contains(10))

	fmt.Println("get", ss.Get(1))

}
