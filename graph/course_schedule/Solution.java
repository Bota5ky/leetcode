package graph.course_schedule;

import java.util.ArrayList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/course-schedule/">207. 课程表</a>
 * @since 2023-09-14 16:22
 */
class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        // [a, b]  b->a
        ArrayList<Integer>[] table = new ArrayList[numCourses];
        for (int i = 0; i < numCourses; i++) {
            table[i] = new ArrayList<>();
        }
        int[] inDegrees = new int[numCourses];
        for (int[] p : prerequisites) {
            table[p[1]].add(p[0]);
            ++inDegrees[p[0]];
        }

        int remain = numCourses;
        for (; ; ) {
            int preRemain = remain;
            for (int i = 0; i < inDegrees.length; i++) {
                if (inDegrees[i] == 0) {
                    --inDegrees[i];
                    remain--;
                    ArrayList<Integer> nextCourses = table[i];
                    for (Integer nextCourse : nextCourses) {
                        --inDegrees[nextCourse];
                    }
                }
            }
            if (preRemain == remain) {
                break;
            }
        }
        return remain != 0;
    }
}
