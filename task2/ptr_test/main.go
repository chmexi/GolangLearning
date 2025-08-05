package main

import (
	"fmt"
)

func ptr_test(pInt *int) {
	*pInt += 10
}

func main() {
	a := 1
	fmt.Println("a before: ", a)
	ptr_test(&a)
	fmt.Println("a after : ", a)
	pA := &a
	fmt.Println("pA before: ", *pA)
	ptr_test(pA)
	fmt.Println("pA after : ", *pA)
}
