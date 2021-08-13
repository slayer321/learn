package DSA_using_java.array;

import java.util.*;

class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int i = 0, j = 0;
        // List<Integer> res = new ArrayList<Integer>();
        int n = nums.length;
        int res[] = new int[n - k + 1];
        Deque<Integer> q = new ArrayDeque<Integer>();

        while (j < n) {
            while (q.size() > 0 && q.getLast() < nums[j]) {
                q.pop();
            }
            q.add(nums[j]);

            if (j - i + 1 < k) {
                j++;
            } else if (j - i + 1 == k) {
                res[i] = q.getFirst();

                if (q.getFirst() == nums[i]) {
                    q.poll();
                }
                i++;
                j++;
            }

        }

        return res;

    }

    public static void main(String[] args) {
        int[] nums = { 1, 3, 1, 2, 0, 5 };
        int k = 3;
        Solution obj = new Solution();
        System.out.println((obj.maxSlidingWindow(nums, k)));

    }
}