def maxSubArray(self, nums: List[int]) -> int:
    sum_ = 0
    max_val = nums[0]
    for i in range(len(nums)):
        sum_ += nums[i]
        # print(sum_)
        if sum_ > max_val:
            max_val = sum_
        if sum_ < 0:
            sum_ = 0
    return max_val
