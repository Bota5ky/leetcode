package sort.queue_reconstruction_by_height;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/queue-reconstruction-by-height/">406. 根据身高重建队列</a>
 * @since 2023-09-07 13:56
 */
class Solution {
    public int[][] reconstructQueue(int[][] people) {
        //先按照个子从高到低排序，如果个子一样，则按照k从小到大排列，这样就得到了一个方便后面插入的队列
        Arrays.sort(people, (o1, o2) -> {
            if (o1[0] != o2[0]) {
                return o2[0] - o1[0];
            } else {
                return o1[1] - o2[1];
            }
        });
        for (int i = 0; i < people.length; i++) {
            int[] temp = new int[2];
            temp[0] = people[i][0];
            temp[1] = people[i][1];
            for (int j = i - 1; j >= temp[1]; j--) {
                people[j + 1] = people[j];
            }
            people[temp[1]] = temp;
        }
        return people;
    }
}
