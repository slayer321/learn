# def binarySearch(arr, l, r, n):
#     while l <= r:
#         mid = l + (r - l) // 2
#         if arr[mid] == n:
#             return arr[mid]
#         if arr[mid] > n:
#             r = mid - 1
#         elif arr[mid] < n:
#             l = mid + 1

#     return -1


# recursive

def binarySearch(arr, l, r, n):
    while l <= r:
        mid = l + (r - l) // 2
        if arr[mid] == n:
            return arr[mid]
        if arr[mid] > n:
            return binarySearch(arr, l, mid - 1, n)
        elif arr[mid] < n:
            return binarySearch(arr, mid + 1, r, n)

    return -1


arr = [1, 2, 3, 4, 5, 6, 7, 21, 43, 65, 77, 88]
l = 0
r = len(arr) - 1
n = 88
print(binarySearch(arr, l, r, n))
