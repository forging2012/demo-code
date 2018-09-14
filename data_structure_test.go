package demo_code

import (
	"fmt"
	"stest/demo-code/utils"
	"testing"
)

// test arithmetic
func TestStringBFMethod(t *testing.T) {
	p := "hello"
	s := "ll"
	sindex := utils.BFArithmetic([]byte(p), []byte(s))
	fmt.Println("localtion=", sindex)
}

// 测试爬楼梯问题
func TestCalStep(t *testing.T) {
	n := utils.CalStepOneOrTwo(4)
	fmt.Println(n)
}
