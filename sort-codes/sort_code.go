##基本的排序代码和理解
##create 2018.02.23


分治思想：
	　将原问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题，然后将这些子问题的解组合为原问题的解。

data := []int{1, 3, 2, 8, 5}
一、排序
	1、冒泡排序		
			冒泡排序：
				内部比较排序
			内容：
				 外循环控制循环次数
				 内部循环用来两个相邻的数交换位置
		
		func BubbleSort(data []int) {
			for i := 0; i < len(data)-1; i++ { // 控制循环次数：i外循环的循环次数
				for j := 0; j < len(data)-i-1; j++ { // 判断每一次循环的时候,相邻的两个数字用来交换位置
					if data[j+1] < data[j] {
						data[j], data[j+1] = data[j+1], data[j]
					}
				}
			}
		}
2、快速排序
	首先：从一堆杂乱的数组中挑出一个基准元素, 
	接着：将这个数组分区，比这个基准小的放在左边,比这个基准大的放在右边。
	最后：重复上面这两个步骤  00

二、查找
	1、二分查找
	印象中的二分查找的基础好像是对于已经排序好的数组。
	// 二分查找
	func BinarySearch(n int, data []int) (int) {
		min := 0
		max := len(data) - 1

		for min <= max {
			mid := (min + max) / 2
			if n > data[mid] {
				min = mid + 1
			} else if n < data[mid] {
				max = mid - 1
			} else {
				return mid + 1
			}
		}

		return -1
	}

