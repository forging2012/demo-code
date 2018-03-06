## 查找代码
## 2018.3.6

func Test_Find_Element(T *testing.T) {

	array := []int{5, 3, 8, 0, 1}
	array = select_sort(array, len(array))
	n := 3
	//index := binary_search(array, n)
	//index := sequence_search(array, n)
	index := difference_search(array, n)
	fmt.Println(array)
	fmt.Println(index)
}

// 差值查找---二分查找的改进版本,每次并不一定是要求折半
func difference_search(array []int, n int) int {

	low := 0
	high := len(array) - 1

	for {
		rate := (n - array[low]) / (array[high] - array[low])

		if rate < 0 || rate > 1 { // 查找的比例有问题
			return -1
		}
		mid := low + (high-low)*rate

		if n > array[mid] {
			low = mid + 1
		} else if n < array[mid] {
			high = high - 1
		} else {
			return mid
		}

	}

	return -1
}

// 二分查找----也叫折半查找
func binary_search(array []int, n int) int {

	min := 0
	max := len(array) - 1

	for {
		mid := (min + max) / 2

		if n > array[mid] {
			min = mid + 1
		} else if n < array[mid] {
			max = mid - 1
		} else {
			return mid
		}
	}

	return -1

}

// 顺序查找
func sequence_search(array []int, n int) int {

	for i := 0; i < n; i++ {
		if array[i] == n {
			return i
		}
	}

	return -1
}

/*

	简单排序
*/

// 插入排序--内部稳定排序
func insert_sort(array []int, n int) []int {

	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}
	return array
}

// 选择排序--内部不稳定排序
func select_sort(array []int, n int) []int {

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}

	return array
}

// 冒泡排序--内部不稳定排序
func bubble_sort(array []int, n int) []int {

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array

}
