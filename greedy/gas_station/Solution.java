package greedy.gas_station;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/gas-station/">134. 加油站</a>
 * @since 2023-09-18 15:27
 */
class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int sum = 0;
        int min = 0;
        int idx = 0;
        for (int i = 0; i < gas.length; i++) {
            sum += gas[i] - cost[i];
            if (sum < min) {
                min = sum;
                idx = i + 1;
            }
        }
        return sum < 0 ? -1 : idx;
    }
}
