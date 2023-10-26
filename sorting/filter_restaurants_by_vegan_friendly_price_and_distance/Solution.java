package sorting.filter_restaurants_by_vegan_friendly_price_and_distance;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/filter-restaurants-by-vegan-friendly-price-and-distance/">1333. 餐厅过滤器</a>
 * @since 2023-09-27 09:48
 */
class Solution {
    public List<Integer> filterRestaurants(int[][] restaurants, int veganFriendly, int maxPrice, int maxDistance) {
        var priorQueue = new PriorityQueue<>((Comparator<int[]>) (r1, r2) -> r1[1] != r2[1] ? r2[1] - r1[1] : r2[0] - r1[0]);
        for (int[] r : restaurants) {
            if (r[2] >= veganFriendly && maxPrice >= r[3] && maxDistance >= r[4]) {
                priorQueue.offer(r);
            }
        }
        var result = new ArrayList<Integer>(priorQueue.size());
        while (!priorQueue.isEmpty()) {
            result.add(priorQueue.poll()[0]);
        }
        return result;
    }
}
