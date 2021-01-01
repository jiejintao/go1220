package util

import (
	"fmt"
	"testing"
)

func TestCRead(t *testing.T) {
	fmt.Println("请输入一个值")
	c := CRead()
	fmt.Println("值: ", c)
	t.Log(c)
}