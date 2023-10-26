package shuffle_an_array

import "math/rand"

// Solution 384. 打乱数组 https://leetcode.cn/problems/shuffle-an-array/
type Solution struct {
	Recover []int
	Nums    []int
}

func Constructor(nums []int) Solution {
	var ret Solution
	ret.Recover = make([]int, len(nums))
	copy(ret.Recover, nums)
	ret.Nums = nums
	return ret
}

/** Resets the array to its original configuration and return it. */
func (t *Solution) reset() []int {
	copy(t.Nums, t.Recover)
	return t.Nums
}

/** Returns a random shuffling of the array. */
func (t *Solution) shuffle() []int {
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
