def subset(nums):
    out = [[]]
    for num in nums:
        out += [curr + [num] for curr in out]
    return out


nums = [1, 2, 3]
print(subset(nums))
