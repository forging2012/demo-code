package other_code

import (
	"fmt"
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	data := []int{3, 2, 4}
	target := 6
	startTime := time.Now().UnixNano() / 1e6
	result := toSum4(data, target)
	endTime := time.Now().UnixNano() / 1e6
	fmt.Println("target=", target, "result=", result, "subTime=", endTime-startTime)
}
