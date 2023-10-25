package graph.course_schedule_iii;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/course-schedule-iii/">630. 课程表 III</a>
 * @since 2023-09-15 14:54
 */
class Solution {
    public int scheduleCourse(int[][] courses) {
        Arrays.sort(courses, Comparator.comparingInt(a -> a[1]));
        int total = 0;

        var priorQueue = new PriorityQueue<Integer>((a, b) -> (b - a));

        for (int[] course : courses) {
            var costTime = course[0];
            var endDate = course[1];
            if (costTime + total <= endDate) {
                total += costTime;
                priorQueue.offer(costTime);
            } else {
                if (!priorQueue.isEmpty() && costTime < priorQueue.peek()) {
                    total -= priorQueue.poll() - costTime;
                    priorQueue.offer(costTime);
                }
            }
        }

        return priorQueue.size();
    }
}
