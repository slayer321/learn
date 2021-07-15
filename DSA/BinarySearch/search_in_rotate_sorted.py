def search(self, nums: List[int], target: int) -> int:
    # nums = sorted(nums)
    # print(nums)
    left = 0
    right = len(nums) - 1
    while left <= right:
        mid_ixd = (left + right) // 2
        # mid_num = nums[mid_ixd]
        if nums[mid_ixd] == target:
            return mid_ixd
        if nums[mid_ixd] >= nums[left]:
            if nums[left] <= target < nums[mid_ixd]:
                right = mid_ixd - 1
            else:
                left = mid_ixd + 1
        else:
            if nums[mid_ixd] < target <= nums[right]:
                left = mid_ixd + 1
            else:
                right = mid_ixd - 1
    return -1
