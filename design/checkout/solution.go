package dui_lie_de_zui_da_zhi_lcof

// Checkout LCR 184. 设计自助结算系统 https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/
type Checkout struct {
	q       []int
	support []int // 双向队列 值为q中对应最大值的坐标
}

func Constructor() Checkout {
	return Checkout{}
}

func (t *Checkout) Maxvalue() int {
	if len(t.q) == 0 {
		return -1
	}
	return t.q[t.support[0]]
}

func (t *Checkout) Pushback(value int) {
	t.q = append(t.q, value)
	//删除support中所有小于value的值
	for len(t.support) > 0 && t.q[t.support[len(t.support)-1]] < value {
		t.support = t.support[:len(t.support)-1]
	}
	if len(t.support) == 0 || len(t.support) > 0 && t.q[t.support[len(t.support)-1]] >= value {
		t.support = append(t.support, len(t.q)-1)
	}
}

func (t *Checkout) Popfront() int {
	if len(t.q) == 0 {
		return -1
	}
	res := t.q[0]
	t.q = t.q[1:]
	if t.support[0] == 0 {
		t.support = t.support[1:]
	}
	for key := range t.support {
		t.support[key]--
	}
	return res
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
