package gosolutions

import "math/rand"

// https://leetcode-cn.com/problems/shuffle-an-array/
type solution struct {
	Recover []int
	Nums    []int
}

func constructor2(nums []int) solution {
	var ret solution
	ret.Recover = make([]int, len(nums))
	copy(ret.Recover, nums)
	ret.Nums = nums
	return ret
}

/** Resets the array to its original configuration and return it. */
func (t *solution) reset() []int {
	copy(t.Nums, t.Recover)
	return t.Nums
}

/** Returns a random shuffling of the array. */
func (t *solution) shuffle() []int {
	for i := 0; i < len(t.Nums); i++ {
		p := rand.Intn(len(t.Nums)-i) + i
		t.Nums[i], t.Nums[p] = t.Nums[p], t.Nums[i]
	}
	return t.Nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
