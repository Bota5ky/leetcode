package all_oone_data_structure

// AllOne 432. 全 O(1) 的数据结构 https://leetcode.cn/problems/all-oone-data-structure/
type AllOne struct {
	Map  map[string]*Node
	Data []*Node
}

type Node struct {
	Key string
	Val int
	Pos int
}

func Constructor() AllOne {
	return AllOne{
		Map: make(map[string]*Node, 100),
	}
}

func (t *AllOne) Inc(key string) {
	node, ok := t.Map[key]
	if !ok {
		node := &Node{Key: key, Val: 1, Pos: len(t.Data)}
		t.Data = append(t.Data, node)
		t.Map[key] = node
		return
	}
	node.Val++
	for i := node.Pos - 1; i >= 0; i-- {
		if t.Data[i+1].Val > t.Data[i].Val {
			t.Data[i+1], t.Data[i] = t.Data[i], t.Data[i+1]
			t.Data[i+1].Pos++
			t.Data[i].Pos--
		} else {
			break
		}
	}
}

func (t *AllOne) Dec(key string) {
	node, ok := t.Map[key]
	if !ok {
		return
	}
	node.Val--
	for i := node.Pos; i < len(t.Data)-1; i++ {
		if t.Data[i+1].Val > t.Data[i].Val {
			t.Data[i+1], t.Data[i] = t.Data[i], t.Data[i+1]
			t.Data[i+1].Pos++
			t.Data[i].Pos--
		} else {
			break
		}
	}
	lastNode := t.Data[len(t.Data)-1]
	if lastNode.Val == 0 {
		delete(t.Map, lastNode.Key)
		t.Data = t.Data[:len(t.Data)-1]
	}
}

func (t *AllOne) GetMaxKey() string {
	if len(t.Data) == 0 {
		return ""
	}
	return t.Data[0].Key
}

func (t *AllOne) GetMinKey() string {
	if len(t.Data) == 0 {
		return ""
	}
	return t.Data[len(t.Data)-1].Key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
