def solve(nums):
    ls = 0
    num_set = set(nums)
    for num in num_set:
        if num - 1 not in num_set:
            curren_num = num
            curr_s = 1
            while curren_num+1 in num_set:
                curren_num += 1
                curr_s += 1
            ls = max(curr_s, ls)
    return ls


nums = [100, 4, 200, 1, 3, 2]
print(solve(nums))
