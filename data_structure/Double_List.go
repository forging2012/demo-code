// 双向链表--golang 实现
// ## 2018.3.8

// 参考网址：https://segmentfault.com/a/1190000011929534


package algorithm_code_self


type Node struct {
	previous *Node
	data     interface{}
	next     *Node
}

type DoubleList struct {
	head  *Node //头结点
	tail  *Node // 尾节点
	count int   // 链表的长度
}

func (d *DoubleList) Append(v interface{}) {
	e := &Node{data: v} // 创建新的节点

	if d.head == nil { // 表示是一个空的双向链表
		d.head = e // 链表的头尾归属
		d.tail = e

		e.previous = nil // 这个新增的头街点的前继和后继节点指向均为nil
		e.next = nil     // 新增节点的后继节点不存在,所以为nil
	} else {

		e.previous = d.tail // 新增节点指向最后一个尾部节点
		e.next = nil

		d.tail.next = e // 让原先链表的最后一个节点指向新增节点
		d.tail = e      // 移动链表的指针,让其指向新增加的元素

	}

	d.count++
}

func (d *DoubleList) InsertBefore(e *Node, v interface{}) bool {

	if e == nil {
		return false
	}

	if d.isHead(e) {
		n := &Node{data: v}
		n.next = d.GetHead() // 新增节点是头结点,现在指向未插入之前的头结点
		n.previous = nil     // 新的头结点的前继节点为空

		d.head.previous = n // 旧的头结点指向新的头结点
		d.head = n          // 头指针指向新的头结点
		d.count++
		return true
	} else { // 直接调用已经写好的方法进行追加节点就好
		return d.InsertAfter(e.previous, v)
	}

	return true
}

func (d *DoubleList) InsertAfter(e *Node, v interface{}) bool {

	if e == nil {
		return false
	}

	if d.isTail(e) { // 尾节点,直接追加
		d.Append(v)
	} else {
		//  node1--->node2--->node3--->node4
		// 比如在node2和node4之间插入node3
		n := &Node{data: v}
		n.next = e.next     // 新增加节点指向插入节点的下一个节点  :  node3指向node4
		e.next.previous = n // 新增加节点的下一个节点指向新增加节点： node4指向node3

		e.next = n     // 新增加节点的上一个节点指向新增加节点：node2指向node3
		n.previous = e // 新增加节点指向新增加节点的上一个节点
		d.count++      // 链表的数量加1
	}

	return true
}

//
func (d *DoubleList) DeleteNode(e *Node) interface{} {

	if e == nil {
		return nil
	}
	//  node1--->node2--->node3--->node4
	switch {
	case d.isHead(e): // 这时候删除node1
		d.head = e.next            // 改变头指针的位置
		d.GetHead().previous = nil // 头结点的前继节点为空
	case d.isTail(e):
		d.tail = e.previous
		d.tail.next = nil
	default: // 比如删除node3
		next := e.next
		previous := e.previous

		previous.next = next     // node2指向node4
		next.previous = previous // node4指向node2
	}

	d.count--
	return e.GetData()
}

// isHead:判断是否是头结点
func (d *DoubleList) isHead(v *Node) bool {
	return d.GetHead() == v
}

// isTail:判断是否是头结点
func (d *DoubleList) isTail(v *Node) bool {

	return d.GetTail() == v
}

// 获取某个节点的数据
func (n *Node) GetData() interface{} {
	return n.data
}

// GetHead:返回列表的头部节点
func (d *DoubleList) GetHead() *Node {
	return d.head
}

// GetTail:返回列表的尾节点
func (d *DoubleList) GetTail() *Node {
	return d.tail
}

// Count返回列表的长度
func (d *DoubleList) Count() int {

	if d == nil {
		return 0
	}
	return d.count
}

// New创建一个新的链表
func (d *DoubleList) New() *DoubleList {
	newList := &DoubleList{}
	newList.count = 0
	newList.head = nil
	newList.tail = nil
	return newList
}
