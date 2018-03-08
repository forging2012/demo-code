##BF算法
##2018.3.8

模式匹配的内容
	模式匹配最朴素的算法是回溯法，
	即模式串跟主串一个字符一个字符的匹配，当模式串中跟主串不匹配时，
	主串回溯到与模式串匹配开始的下一个位置，模式串回溯到第一个位置，继续匹配。算法的时间复杂度为O（m*n）

bf算法每次从主串的第一个位置开始，找不到不匹配的，模式串就向后移动以为，无脑式操作。
在解决实际问题的时候,忧郁数据量庞大,时间复杂度往往会很高

func Test_bf(T *testing.T) {
	p := "hello"
	s := "ll"
	fmt.Println(Index([]byte(p), []byte(s)))
}


// Big:是主串,Small:是模式串
func Index(Big, Small []byte) int {
	fmt.Println(len(Big), len(Small))

	i, j := 0, 0

	for i < len(Big) && j < len(Small) {
		if Big[i] == Small[j] {
			i++
			j++
		} else {
			i = i - j + 1 //如果有一位没有匹配上,这时候就主串就向后移动以为,模式穿继续从0开始匹配
			j = 0
		}
	}

	if j == len(Small) { // 模式串和主串相匹配
		return i - len(Small)
	}

	return -1 // 表示没有找到与之匹配的
}
