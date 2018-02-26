##基本的排序代码和理解
##create 2018.02.23

排序：
	冒泡排序
	快速排序
	归并排序

稳定排序
	插入排序(未实现)、冒泡排序、归并排序
不稳定排序
	未实现：
		快速排序、	
	已实现：
		选择排序、希尔排序、堆排序
参考网址：
	http://www.cnblogs.com/UnGeek/p/5645534.html

排序算法的稳定性
	通俗地讲就是能保证排序前2个相等的数其在序列的前后位置顺序和排序后它们两个的前后位置顺序相同。
	在简单形式化一下，如果Ai = Aj，Ai原来在位置前，排序后Ai还是要在Aj位置前。
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
2、快速排序(不稳定排序)
	
	介绍快速排序思想的一个网址,很容易理解：
		http://www.cnblogs.com/morewindows/archive/2011/08/13/2137415.html.
	根据这个网址,快速排序就是挖坑填数+分治法
		找到合适的数就填入到上一个挖好的坑中，比如刚刚开始，挖好的坑就是最左边的那个数
	首先：从一堆杂乱的数组中挑出一个基准元素, 
	接着：将这个数组分区，比这个基准小的放在左边,比这个基准大的放在右边。
	最后：重复上面这两个步骤  00
	func QuickSort(data []int, start, end int) {
		key := data[start] // 标准元素
		low := start
		high := end

		for {

			for low < high {
				if data[high] < key {
					data[low] = data[high]
					break
				}
				high--
			}

			for low < high {
				if data[low] > key {
					data[high] = data[low]
					break
				}
				low++
			}

			if low >= high {
				data[low] = key
				break
			}
		}

		if low-1 > start {
			QuickSort(data, start, low-1)
		}

		if high+1 < end {
			QuickSort(data, high+1, end)
		}
	}

3、归并排序(稳定排序)
	 参考网址：
		http://www.cnblogs.com/skywang12345/p/3602369.html
		从上往下的归并排序：它与"从下往上"在排序上是反方向的。它基本包括3步：
		① 分解 -- 将当前区间一分为二，即求分裂点 mid = (low + high)/2;
		② 求解 -- 递归地对两个子区间a[low...mid] 和 a[mid+1...high]进行归并排序。递归的终结条件是子区间长度为1。
		③ 合并 -- 将已排序的两个子区间a[low...mid]和 a[mid+1...high]归并为一个有序的区间a[low...high]。

理解：
	就是不等的分成两部分切片。直至左边的切片和右边的切片长度都为1，然后返回。
程序把返回的左切片和右切片进行排序。不停地重复。
也就是一个类似汉诺塔的问题。
		func merge_sort(array []int) []int {
			l := len(array) // 如果数组长度为1,递归结束,否则再次将数组等分(也就是判断递归是否结束)
			if l == 1 {
				return array
			}

			mid := len(array) / 2
			left := merge_sort(array[0:mid])
			right := merge_sort(array[mid:])

			return sort(left, right)
		}

		// 自上而下进行排序
		func sort(left, right []int) (result []int) {

			i, j := 0, 0

			// 保证长度最小的数组全部添加到这个result切片中
			for i < len(left) && j < len(right) {
				if left[i] < right[j] {
					result = append(result, left[i])
					i++
				} else {
					result = append(result, right[j])
					j++
				}
			}
			
			// 将长度较长的已经排序好的数组添加到这个result数组中
			
			result = append(result, left[i:]...)
			result = append(result, right[j:]...)

			return
		}

4、插入排序(稳定排序)
	定义：插入排序就是每一步都将一个待排数据按其大小插入到已经排序的数据中的适当位置，直到全部插入完毕。
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

