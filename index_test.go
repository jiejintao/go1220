package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T)  {
	a := 1
	b := 2
	c := add(a,b)
	fmt.Println(c)
	t.Log(c)
}
