package sort_code_review

import (
	"fmt"
	"sort"
	"testing"
)

func TestHello(t *testing.T) {
	array := []int{6, 4, 8, 1, 0, 2}

	//fmt.Println(BinarySearch(array, 100))

	//bubble_sort(array)
	//quick_sort(array, 0, len(array)-1)
	//insert_sort(array)
	//select_sort(array)

	//fmt.Println(array)
	fmt.Println(merge_sort(array))
}

/*

归并排序:---递归
	就是将数组不停地划分,最后对两个数组进行排序!
*/

func merge_sort(data []int) (result []int) {
	if len(data) == 1 {
		result = data
		return
	}

	mid := len(data) / 2
	left := merge_sort(data[:mid])
	right := merge_sort(data[mid:])
	result = sort_arrays(left, right)
	return
}

func sort_arrays(left, right []int) (result []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) { // 将两个数组逐个进行比较,将每个数组中较小的,添加到这个result数组中
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...) // 将剩余的数组中的元素,追加到result数组中
	return
}

// 二分查找,前提是已经排好序的数组
func BinarySearch(data []int, num int) (index int) {
	sort.Ints(data)
	min := 0
	max := len(data) - 1
	for min <= max {
		mid := (min + max) / 2

		if num > data[mid] {
			min = mid + 1
		} else if num < data[mid] {
			max = mid - 1
		} else {
			index = mid
			break
		}
	}
	return -1
}

/*
选择排序:就是从当前数组中找出最小的，然后放到第一位,
接着找出次小的放到第二位
*/
func select_sort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}

		if min != i {
			data[min], data[i] = data[i], data[min]
		}
	}
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
