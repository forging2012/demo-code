package utils

/*
走台阶:总共n级台阶,一次走一级,一次走两级,总共有多少级算法
*/

func CalStepOneOrTwo(n int) int {
	if n == 0 || n == 1 || n == 2 || n == 3 {
		if n == 3 {
			return n + 1
		}
		return n
	} else {
		return CalStepOneOrTwo(n-1) + CalStepOneOrTwo(n-2) + CalStepOneOrTwo(n-3)
	}
}
