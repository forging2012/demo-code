##基本的排序代码和理解
##create 2018.02.23

稳定排序
	插入排序、冒泡排序、归并排序
不稳定排序
	已实现：
		快速排序、选择排序、希尔排序
	未实现：
		堆排序
参考网址：
	http://www.cnblogs.com/UnGeek/p/5645534.html

维基百科：
	https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8Fhttps://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F
只需要更换最后几个排序的名字就可以查看其它排序方法

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

4、插入排序(稳定排序--内部比较)
	定义：插入排序就是每一步都将一个待排数据按其大小插入到已经排序的数据中的适当位置，直到全部插入完毕。
	自己的理解：也就是从这个数组的第二个元素开始。
		依次和前面的那个数进行比较。比前边的数大的话,就原地不动。
								比前边的数小，这两个数就交换位置(将前面的数往后移动一位)
	func insert_sort(data []int) []int {

		for j := 1; j < len(data); j++ {

			i := j
			for i-1 >= 0 && data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
				i--
			}

		}
		return data
	}

	//  推荐这种方式
	func insert_sort2(data []int) []int {

		for i := 1; i < len(data); i++ {
			key := data[i]
			j := i - 1

			for j >= 0 && data[j] > key {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = key

		}

		return data
	}
5、选择排序(内部排序--不稳定排序)
	定义：遍历数组，依次从数组中找出最小的元素,然后与第一个元素进行交换。依次类推
	流程:
		也就是每次循环的时候，假设第一个元素是最小元素，从第二个到最后一个，在这个范围内找出最小的元素。
	找到之后,然后和第一个进行交换。（每次找到一个最小元素，就把被找到的元素和最小元素的坐标互换）

	// 选择排序
	func test_select_sort(array []int, n int) []int {
		for i := 0; i < n-1; i++ { // 外循环控制循环次数
			index := i                   // 记录最小元素的索引
			for j := i + 1; j < n; j++ { // 内循环用于找出最小元素,并和第一个进行交换
				if array[j] < array[index] {
					index = j
				}
			}
			// 判断找到最小元素的索引和找到数组中的索引是否相同
			if index != i {
				array[i], array[index] = array[index], array[i]
			}
		}
		return array
	}
6、希尔排序(内部排序--不稳定排序)
	定义：将无序数组分割为若干个子序列，子序列不是逐段分割的，而是相隔特定的增量的子序列，对各个子序列进行插入排序；
		然后再选择一个更小的增量，再将数组分割为多个子序列进行排序......
		最后选择增量为1，即使用直接插入排序，使最终数组成为有序。
	参考网址：
		http://www.cnblogs.com/skywang12345/p/3597597.html#a1--这个分开写了
	/*
		one:也就是按照一定的步长进行分组,然后对每一个分组进行插入排序。
		two:重复上述步骤。
		three:最后以步长为1进行插入排序
	*/

	// data:待排序的数组
	// n:待排序的数组长度
	func shell_sort1(data []int, n int) []int {

		var i, step int
		for step = n / 2; step > 0; step /= 2 { // 总共有几个step
			// 单独对某一组进行排序
			for i = 0; i < step; i++ { // 这里也是我之前没有用到过,把99个数分成两组,这样每一组都有49个数,最多多的数用步长为1进行插入排序
				group_srot(data, n, i, step)
			}
		}
		return data
	}

	/*
		对同一个分组内元素进行插入排序
		array:待排序的数组
		n:数组的长度
		step:表示步长(每次间隔的距离)
		i:代表数组的起始长度
	*/
	func group_srot(array []int, n, i, step int) {

		// 对同一组内的元素进行排序,内循环是对某个数进行排序。外层循环移动数组的下一个元素
		for j := i + step; j < n; j += step {
			if array[j] < array[j-step] { // 这个判断可有可无,不影响程序的运行.但是多了这一个判断.对于tmp,k,array[k+step]的赋值,避免无用的赋值,即array[j]>array[j-step]的情况
				tmp := array[j]
				k := j - step
				for k >= 0 && array[k] > tmp {
					array[k+step] = array[k]
					k -= step
				}
				array[k+step] = tmp
			}
		}
	}
	/*
		这是基于第一种方式的,每次直接从每一组的第二个元素开始排序
		参考网址：http://blog.csdn.net/morewindows/article/details/6668714
	*/
	func shellSort2(array []int, n int) []int {
		var j, step int

		for step = n / 2; step > 0; step /= 2 { // 首先分组
			for j = step; j < n; j++ { // 从数组的第step个元素开始
				if array[j]<array[j-step]{ // 每个元素与自己数组内的数据进行直接插入排序---(这半部分就是直接插入排序)
					tmp := array[j]
					k := j - step
					for k >= 0 && array[k] > tmp {
						array[k+step] = array[k]
						k -= step
					}
					array[k+step] = tmp
				}
			}
		}

		return array
	}

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

