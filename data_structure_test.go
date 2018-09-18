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

func TestKMP(t *testing.T) {
	//content := []byte("what's your name? and Todaya isi Friday")
	content := []byte("what's aaaaa and Todaya isi Friday")
	sub := []byte("hat")

	index, _ := KMP(content, 0, len(content)-1, sub)
	fmt.Println("final=", index)

}

func GetNextValueArray(sub []byte) (next []int) {
	fmt.Println("sub=", sub)
	var (
		length        int = len(sub)
		middle        int
		compare_left  int
		compare_right int
		match_count   int
	)

	next = make([]int, length)
	next[0] = 0
	next[1] = 0 // next[2]=1   next[3]=0   next[4]=1
	// aa b aa   // [0 0 1 2 2]
	for i := 2; i < length; i++ {
		middle = i / 2
		match_count = 0

		tmp := i % 2
		if tmp == 0 {
			for j := 0; j < middle; { //i=2,middle=1   j=0,
				compare_left = 0          //compare_left=0
				compare_right = i - 1 - j //compare_right=1
				for compare_left <= j {   //match_count=
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left++
					compare_right++
				}
				if compare_left == j+1 { // 保证最大连续子串相等
					match_count++
				}
				j++
			}
			next[i] = match_count
		} else { // i=3的时候,第一次是不相等,第二次是compare_left!=j+1,所以match_count=1
			for j := 0; j <= middle; j++ { // middle=1  i=3   j=0
				compare_left = 0          // compare_left=0   1  2
				compare_right = i - 1 - j // compare_right=2  3  4
				for compare_left <= j {   //  0<=2  match_count=
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left++
					compare_right++
				}
				if compare_left == j+1 {
					match_count++
				}
			}
			next[i] = match_count
		}
	}
	return next
}

//在content中的start-end之间寻找sub子串
//成功返回匹配成功的起始下标，匹配失败则返回-1
func KMP(content []byte, start_index int, end_index int, sub []byte) (index int, cc []int) {

	cc = GetNextValueArray(sub)
	fmt.Println("cc=", cc)
	var (
		next       []int = ReviseNextValueArray(cc)
		sub_index  int   = 0
		sub_length int   = len(sub)
	)
	fmt.Println("next=", next)
	//return
	for i := start_index; i <= end_index; i++ {
		if content[i] == sub[sub_index] {
			match_start := i
			for j := sub_index; j <= sub_length; j++ {
				if j == sub_length {
					return match_start - sub_index, cc
				}
				if i >= end_index || content[i] != sub[j] {
					sub_index = next[j]
					break
				}
				i++
			}
		}
	}

	return -1, cc
}

func ReviseNextValueArray(next []int) []int {
	length := len(next)
	for i := 2; i < length; i++ {
		if next[i] == next[next[i]] {
			next[i] = next[next[i]]
		}
	}
	return next
}
