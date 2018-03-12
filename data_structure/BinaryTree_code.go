## 简单的平衡二叉树
## 2018.3.12
/*
	二叉树
	create at 2018.2.12
*/

参考网址:
		http://blog.csdn.net/skh2015java/article/details/79242356

type BinaryTree struct {
	data  int         // 代表当前节点存储的值
	left  *BinaryTree // 代表左子树
	right *BinaryTree // 代表右子树
}

// CreateBinaryTree:创建一个新的二叉树---（左子树和右字树都不存在）
func CreateBinaryTree(rootValue int) *BinaryTree {

	return &BinaryTree{data: rootValue}
}

// Insert:向二叉树中添加新元素(感觉包含了递归......)
func (this *BinaryTree) Insert(appendValue int) *BinaryTree {

	if this == nil {
		this = CreateBinaryTree(appendValue)
		return this
	}

	if this.data > appendValue {// 这里隐含了,如果左子树存在的话，将会不停的调用Insert方法,直到左子树不存在就会创造一个节点。
		this.left = this.left.Insert(appendValue)
	} else {
		this.right = this.right.Insert(appendValue)
	}
	return this
}

// Contains:判断二叉树是否包含某一个数
func (this *BinaryTree) Contains(containValue int) bool {

	if this == nil {
		return false
	}

	v := this.comapreTo(containValue)

	if v < 0 { // 判断左边是否包含
		return this.left.Contains(containValue)
	} else if v > 0 { // 判断右边是否包含
		return this.right.Contains(containValue)
	} else { // 正好匹配
		return true
	}

}
func (this *BinaryTree) comapreTo(value int) int {
	return this.data - value
}

// Remove:移除元素
func (this *BinaryTree) Remove(value int) *BinaryTree {

	if this == nil {
		return nil
	}

	return this
}

// FindMax:查找最大元素
func (this *BinaryTree) FindMax() int {

	if this == nil {
		fmt.Println("Empty BinaryTree")
		return -1
	}

	if this.right == nil {
		return this.data
	} else {
		return this.right.FindMax()
	}
}

// FindMaxNode:查找最大的节点
func (this *BinaryTree) FindMaxNode() *BinaryTree {

	if this.right != nil {
		for this.right != nil { // 这里利用for循环,不停地查找根节点右边的元素,直到右边的元素不存在,然后跳出循环.
			this = this.right
		}
	}

	return this
}

// FindMin:查找最小元素
func (this *BinaryTree) FindMin() int {

	if this == nil {
		fmt.Println("Empty BinaryTree")
		return -1
	}

	if this.left == nil {
		return this.data
	} else {
		return this.left.FindMin()
	}

}

// FindMinNode:查找最小节点
func (this *BinaryTree) FindMinNode() *BinaryTree {

	if this.left != nil {
		for this.left != nil {
			this = this.left
		}
	}
	return this
}

// GetAlls:获取所有的数据
func (this *BinaryTree) GetAlls() []int {

	if this == nil {
		return []int{}
	}

	return []int{}
}

func appendValue(result []int, tree *BinaryTree) []int {

	if tree != nil {
		result = appendValue(result, tree.left)
		result = append(result, tree.data)
		result = appendValue(result, tree.right)
	}

	return result
}
