# Find First and Last Position of Element in Sorted Array


def firstOcc(nums, k):
    l = 0
    r = len(nums) - 1
    while l <= r:
        mid = l + (r - l)//2
        if nums[mid] == k:
            if mid-1 >= 0 and nums[mid-1] == k:
                r = mid - 1
            else:
                return mid
        elif k < nums[mid]:
            r = mid - 1
        else:
            l = mid + 1
    return -1


def secondOcc(nums, k):
    l = 0
    r = len(nums) - 1
    while l <= r:
        mid = l + (r - l)//2

        if nums[mid] == k:
            if mid+1 < len(nums) and nums[mid+1] == k:
                l = mid + 1
            else:
                return mid
        elif k < nums[mid]:
            r = mid - 1
        else:
            l = mid + 1
    return -1


def searchRange(nums, k):
    if not nums:
        return [-1, -1]
    return [firstOcc(nums, k), secondOcc(nums, k)]


nums = [2, 2]
k = 2
print(searchRange(nums, k))
