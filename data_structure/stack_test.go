## 队列--golang简易实现
## 2018.3.8



type Element struct {
	data interface{} // 数据域
	next *Element    // 下一个节点的位置
}

type Stack struct {
	top   *Element // 栈顶指针
	count int      // 栈的长度
}

// Append 向栈中添加元素
func (s *Stack) Append(v interface{}) bool {

	if v == nil {
		return false
	}

	n := &Element{data: v}

	if s == nil {
		n.next = nil
		s.top = n
	} else {
		n.next = s.top // 新增加的栈顶元素指向之前的栈顶元素
		s.top = n      // 更改栈顶指针的位置
	}

	s.count++
	return true

}

// Delete 删除栈顶元素
func (s *Stack) Delete() bool {
	// 队列为空
	if s == nil {
		return false
	}

	// 改变栈顶指针的位置
	s.top = s.top.next
	s.count--
	return true
}

// GetData 获取栈顶的数据
func (s *Stack) GetData() interface{} {

	if s == nil {
		return nil
	}
	return s.top.data
}

// GetCount:返回了堆栈中的元素数量
func (s *Stack) GetCount() int {
	return s.count
}

