package common.question_1350_to_inf.course_schedule_iv;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/course-schedule-iv/">1462. 课程表 IV</a>
 * @since 2023-09-12 20:29
 */
class Solution {
    public List<Boolean> checkIfPrerequisite(int numCourses, int[][] prerequisites, int[][] queries) {
        List<Integer>[] g = new List[numCourses];
        for (int i = 0; i < numCourses; i++) {
            g[i] = new ArrayList<>();
        }

        var inDegree = new int[numCourses];
        for (int[] p : prerequisites) {
            ++inDegree[p[1]];
            g[p[0]].add(p[1]);
        }
        Queue<Integer> queue = new ArrayDeque<>();
        for (int i = 0; i < numCourses; ++i) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        var isPre = new boolean[numCourses][numCourses];
        while (!queue.isEmpty()) {
            var father = queue.poll();
            var sons = g[father];
            for (Integer son : sons) {
                isPre[father][son] = true;
                --inDegree[son];
                for (int i = 0; i < numCourses; ++i) {
                    isPre[i][son] = isPre[i][son] | isPre[i][father];
                }
                if (inDegree[son] == 0) {
                    queue.offer(son);
                }
            }
        }

        var result = new ArrayList<Boolean>(queries.length);
        for (int[] query : queries) {
            result.add(isPre[query[0]][query[1]]);
        }
        return result;
    }
}
