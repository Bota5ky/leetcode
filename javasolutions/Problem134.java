package javasolutions;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/gas-station/">134. 加油站</a>
 * @since 2023-09-07 11:31
 */
class Problem134 {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        for (int i = 0; i < gas.length; i++) {
            int cnt = 0;
            int j = i;
            for (int fuel = 0; fuel >= 0;) {
                fuel += gas[j] - cost[j];
                if (fuel >= 0) {
                    cnt++;
                    j++;
                    if (j == gas.length) {
                        j = 0;
                    }
                }
                if (cnt == gas.length) {
                    return i;
                }
            }
        }
        return -1;
    }
}
