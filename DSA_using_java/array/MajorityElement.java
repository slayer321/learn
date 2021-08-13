package DSA_using_java.array;

import java.util.*;

// Boyer-Moore Voting Algorithm

public class MajorityElement {
    public List<Integer> majorityElement(int[] nums) {
        int num1 = -1, num2 = -1, c1 = 0, c2 = 0, len = nums.length;

        System.out.println(len / 3);
        for (int j = 0; j < len; j++) {
            if (nums[j] == num1) {
                c1++;
            } else if (nums[j] == num2) {
                c2++;
            } else if (c1 == 0) {
                num1 = nums[j];
                c1++;
            } else if (c2 == 0) {
                num2 = nums[j];
                c2++;
            } else {
                c1--;
                c2--;
            }
        }

        System.out.println("num1 " + num1);
        System.out.println("num2 " + num2);
        List<Integer> res = new ArrayList<Integer>();
        int count1 = 0;
        int count2 = 0;
        for (int i = 0; i < len; i++) {
            if (num1 == nums[i]) {
                count1++;
            } else if (num2 == nums[i]) {
                count2++;
            }
        }

        System.out.println("count1 " + count1);
        System.out.println("count2 " + count2);
        if (count1 > (len / 3)) {
            res.add(num1);
        }
        if (count2 > (len / 3)) {
            res.add(num2);

        }

        return res;

    }

    public static void main(String[] args) {
        int[] nums = { 1, 2 };
        MajorityElement obj = new MajorityElement();
        System.out.println(obj.majorityElement(nums));

    }
}
