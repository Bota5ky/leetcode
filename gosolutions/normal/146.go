package leetcode

//LRUCache 缓存机制
//https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	Map  map[int]*Node
	Cap  int
	Head *Node
	Last *Node
}

//Node 储存节点
type Node struct {
	Key, Val  int
	Pre, Next *Node
}

//Constructor Constructor
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Map:  make(map[int]*Node, capacity),
		Cap:  capacity,
		Head: nil,
		Last: nil,
	}
}

//Get Get
func (t *LRUCache) Get(key int) int {
	//不存在
	node, ok := t.Map[key]
	if !ok {
		return -1
	}
	//存在
	t.Remove(node)
	t.SetHeader(node)
	return node.Val
}

//Put Put
func (t *LRUCache) Put(key int, value int) {
	node, ok := t.Map[key]
	if ok {
		node.Val = value
		t.Remove(node)
		t.SetHeader(node)
	} else {
		if len(t.Map) >= t.Cap {
			delete(t.Map, t.Last.Key)
			//先delete 尾结点remove后会变
			t.Remove(t.Last)
		}
		node := &Node{
			Key:  key,
			Val:  value,
			Pre:  nil,
			Next: nil,
		}
		t.Map[key] = node
		t.SetHeader(node)
	}
}

//Remove Remove
func (t *LRUCache) Remove(node *Node) {
	if node.Pre == nil {
		t.Head = node.Next
	} else {
		node.Pre.Next = node.Next
	}
	if node.Next == nil {
		t.Last = node.Pre
	} else {
		node.Next.Pre = node.Pre
	}
}

//SetHeader SetHeader
func (t *LRUCache) SetHeader(node *Node) {
	//头结点为nil
	node.Pre = nil
	if t.Head == nil {
		t.Head = node
		t.Last = node
	} else { //头结不为空
		t.Head.Pre = node
		node.Next = t.Head
		t.Head = node
	}
}
