package common.question_151_to_200.count_primes;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/count-primes/">204. 计数质数</a>
 * @since 2023-09-07 12:43
 */
class Solution {
    public int countPrimes(int n) {
        int[] isPrime = new int[n];
        Arrays.fill(isPrime, 1);
        int ans = 0;
        for (int i = 2; i < n; ++i) {
            if (isPrime[i] == 1) {
                ans += 1;
                if ((long) i * i < n) {
                    for (int j = i * i; j < n; j += i) {
                        isPrime[j] = 0;
                    }
                }
            }
        }
        return ans;
    }
}
