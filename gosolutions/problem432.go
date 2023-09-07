package gosolutions

// AllOne AllOne
// https://leetcode-cn.com/problems/all-oone-data-structure/
type AllOne struct {
	Map  map[string]*Node1
	Data []*Node1
}

// Node1 Node1
type Node1 struct {
	Key string
	Val int
	Pos int
}

// Constructor2 Constructor2
func Constructor2() AllOne {
	return AllOne{
		Map: make(map[string]*Node1, 100),
	}
}

// Inc Inc
func (t *AllOne) Inc(key string) {
	node, ok := t.Map[key]
	if !ok { //不存在
		node := &Node1{Key: key, Val: 1, Pos: len(t.Data)}
		t.Data = append(t.Data, node)
		t.Map[key] = node
		return
	}
	//存在的话
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

// Dec Dec
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
	lastnode := t.Data[len(t.Data)-1]
	if lastnode.Val == 0 {
		delete(t.Map, lastnode.Key)
		t.Data = t.Data[:len(t.Data)-1]
	}
}

// GetMaxKey GetMaxKey
func (t *AllOne) GetMaxKey() string {
	if len(t.Data) == 0 {
		return ""
	}
	return t.Data[0].Key
}

// GetMinKey GetMinKey
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
