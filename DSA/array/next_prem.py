def rev(i, j, arr):
    n = len(arr)
    for k in range(n//2):
        arr[i], arr[j] = arr[j], arr[i]
        i += 1
        j -= 1

    return arr


def nextPermutation(nums):
    """
    Do not return anything, modify nums in-place instead.
    """

    n = len(nums)
    if n <= 1:
        return nums
    i = n-2
    # for i in range(n-2, - 1, -1):
    while i >= 0 and nums[i] >= nums[i+1]:
        i -= 1
    if i == -1:
        return rev(0, n-1, nums)

    for j in range(n-1, -1, -1):
        if nums[j] > nums[i]:
            nums[j], nums[i] = nums[i], nums[j]
            break
    return nums


arr = [1, 2, 3]

print(nextPermutation(arr))
