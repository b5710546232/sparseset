package main

import (
	"fmt"

	"github.com/b5710546232/sparseset/sparseset"
)

func main() {

	// ss := sparseset.NewSparseSet[int](10)
	// ss.Contains(1)
	// ss.Put(1, 99)
	// ss.Contains(1)

	// ss.Put(2, 10)

	// ss.Put(10, 19)

	// ss.Put(7, 7)
	// ss.PrintSpase()
	// ss.PrintDense()
	// ss.Put(7, 999)
	// ss.PrintSpase()
	// ss.PrintDense()
	// ss.Put(7, 99111)
	// ss.PrintSpase()
	// ss.PrintDense()
	// ss.Put(1, 99111)
	// ss.PrintSpase()
	// ss.PrintDense()

	// fmt.Println("------remove------")
	// ss.Remove(2)
	// ss.PrintSpase()
	// ss.PrintDense()

	// ss.PrintSpase()
	// v, ok := ss.Get(1)
	// fmt.Println("get", v, ok)
	// v, ok = ss.Get(5)
	// fmt.Println("get", v, ok)

	ss := sparseset.NewSparseSet[string](1)
	ss.Put(1, "9991")
	ss.Put(2, "21")
	ss.Put(3, "21")
	ss.Put(3, "0")

	val1, _ := ss.Get(1)
	fmt.Println("val1", val1)
}
