package javasolutions;

import java.util.Arrays;
import java.util.Comparator;

//https://leetcode-cn.com/problems/queue-reconstruction-by-height/
class leetcode406 {
    public int[][] reconstructQueue(int[][] people) {
        //先按照个子从高到低排序，如果个子一样，则按照k从小到大排列，这样就得到了一个方便后面插入的队列
        Arrays.sort(people, new Comparator<int[]>() {
            public int compare(int[] o1, int[] o2) {
                if (o1[0] != o2[0]) {
                    return o2[0] - o1[0];
                } else {
                    return o1[1] - o2[1];
                }
            }
        });
        // for (int k = 0; k < people.length; k++) {
        // System.out.printf("%d,%d;", people[k][0], people[k][1]);
        // }
        // System.out.println();
        for (int i = 0; i < people.length; i++) {
            int[] temp = new int[2];
            temp[0] = people[i][0];
            temp[1] = people[i][1];
            for (int j = i - 1; j >= temp[1]; j--) {
                people[j + 1] = people[j];
            }
            people[temp[1]] = temp;
            // for (int k = 0; k < people.length; k++) {
            // System.out.printf("%d,%d;", people[k][0], people[k][1]);
            // }
            // System.out.println();
        }
        return people;
    }

    public static void main(String[] args) {
        int[][] people = { { 7, 0 }, { 4, 4 }, { 7, 1 }, { 5, 0 }, { 6, 1 }, { 5, 2 } };
        //ans [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
        leetcode406 sol = new leetcode406();
        sol.reconstructQueue(people);
    }
}