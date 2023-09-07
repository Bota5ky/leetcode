package javasolutions;

//https://leetcode-cn.com/problems/gas-station/
class leetcode134 {
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