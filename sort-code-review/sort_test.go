package sort_code_review

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	array := []int{6, 4, 8, 1, 0, 2}
	//bubble_sort(array)
	//quick_sort(array, 0, len(array)-1)
	insert_sort(array)
	fmt.Println(array)
}

/*
插入排序
	理解：分为两个数组，假设第一个元素是已经排好序的数组
	从第二个元素开始，依次和前面的进行比较,比之前小的就继续和前边的进行比较
*/
func insert_sort(data []int) {
	for i := 1; i < len(data); i++ { // 从第一个元素开始进行比较
		key := data[i] // 找出标杆元素
		j := i

		for j > 0 && data[j-1] > key {
			data[j] = data[j-1] // 如果前一个元素比较比目标元素小的话,那么就将前一个元素向后移动一位
			j = j - 1
		}
		data[j] = key // 为目标元素找到指定位置
	}
	return
}

/*
	以第一个元素,作为标杆元素,逆向开始,如果这个后面的比前边的小,则交换位置,目标就是
	标杆元素的左边都是比它小的,标杆元素的右边都是比它大的。
	分支法+挖坑填补
*/

// 快速排序
func quick_sort(data []int, beg, end int) {

	low := beg
	high := end
	key := data[beg]

	for {
		for low < high { // 从右边开始找出第一个比标杆元素小的,然后放到之前的位置
			if data[high] < key {
				data[low] = data[high]
				break
			}
			high--
		}
		for low < high { // 从左边开始找出第一个比标杆元素大的元素
			if data[low] > key {
				data[high] = data[low]
				break
			}
			low++
		}

		if low >= high { // 如果两个low和high碰到一起了,则说明标杆元素左边的都是比自己小的,右边的都是比自己大的
			fmt.Println(data)
			data[low] = key
			break
		}
	}

	if low > beg {
		quick_sort(data, beg, low-1)
	}

	if high < end {
		quick_sort(data, low+1, high)
	}
	return
}

/*
冒泡排序:比如升序排序
	理解:就是相邻的两个元素进行比较,大的往后去,
		外循环:控制循环次数
		内循环:就是相邻两个元素进行比较
*/

// 冒泡排序
func bubble_sort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return
}
