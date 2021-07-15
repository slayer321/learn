# Binary search on Reverse Sorted array


def binarysearch(arr, k):
    l = 0
    r = len(arr) - 1
    flag = 0
    if arr[l] < arr[l+1]:
        flag = 0
    else:
        flag = 1

    while l <= r:
        mid = l + (r-l) // 2
        # print(mid)
        if arr[mid] == k:
            return mid
        elif k < arr[mid]:
            l = mid + 1
        elif k > arr[mid]:
            r = mid - 1

        #l += 1
        #r -= 1
    return None


arr = [7, 6, 5, 4, 3, 2, 1]
k = 2
print(binarysearch(arr, k))
