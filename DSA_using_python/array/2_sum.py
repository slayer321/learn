def solve(nums, target):
    nums.sort()
    i = 0
    j = len(nums) - 1
    res = []
    while i < j:
        val = nums[i] + nums[j]
        if val == target:
            res.append(i)
            res.append(j)
            break
        if val > target:
            j -= 1
        else:
            i += 1

    return res


nums = [3, 2, 4]
n = 6
print(solve(nums, n))
