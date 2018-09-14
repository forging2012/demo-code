package utils

func KMPArithmetic() {

}

/*
 bf算法的实现---查找字符串
	模式匹配的内容
		模式匹配最朴素的算法是回溯法，
		即模式串跟主串一个字符一个字符的匹配，当模式串中跟主串不匹配时，
		主串回溯到与模式串匹配开始的下一个位置，模式串回溯到第一个位置，继续匹配。算法的时间复杂度为O（m*n）

	bf算法每次从主串的第一个位置开始，找不到不匹配的，模式串就向后移动以为，无脑式操作。
	在解决实际问题的时候,忧郁数据量庞大,时间复杂度往往会很高
*/
// Big:是主串,Small:是模式串
func BFArithmetic(Big, Small []byte) (index int) {
	i, j, index, lMax, lMin := 0, 0, -1, len(Big), len(Small)
	for i < lMax && j < lMin {
		if Big[i] == Small[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if lMin == j {
		index = i - len(Small)
	}
	return
}
