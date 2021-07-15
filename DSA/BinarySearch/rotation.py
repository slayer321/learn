'''
Number of Times a Sorted array is Rotated
'''


def findKRotation(arr,  n):
    l = 0
    r = n - 1
    while l <= r:
        mid = l + (r - l)//2
        nxt = (mid + 1) % n
        pre = (mid + n - 1) % n

        if (arr[mid] <= arr[pre]) and (arr[mid] <= arr[nxt]):
            return mid

        if arr[mid] <= arr[r]:
            r = mid - 1
        elif arr[l] <= arr[mid]:
            l = mid + 1

    return 0


arr = [66, 72, 79, 86, 95, 104, 106, 110, 119, 123, 124, 129, 132,
       136, 137, 142, 150, 2, 12, 14, 17, 26, 30, 36, 38, 46, 52, 60]
n = 28
print(findKRotation(arr, n))
