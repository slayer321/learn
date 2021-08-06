def search(nums, target):
    n = len(nums)
    l = 0
    r = n - 1

    while l <= r:
        mid = l + (r-l)//2
        if nums[mid] == target:
            return mid
        if nums[mid] <= nums[r]:
            r = mid - 1
        elif nums[l] <= nums[mid]:
            l = mid + 1
    return -1


nums = [4, 5, 6, 7, 0, 1, 2]
k = 0
print(search(nums, k))
