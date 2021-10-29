def binary(arr, val):
    i = 0
    j = len(arr) - 1
    return solve(arr, val, i, j)


def solve(arr, val, i, j):
    while i <= j:
        mid = i + (j-i) // 2
        if arr[mid] == val:
            return mid
        if val < arr[mid]:
            return solve(arr, val, i, mid-1)
        elif val > arr[mid]:
            return solve(arr, val, mid+1, j)
    return -1


arr = [1, 2, 3, 4, 5, 6, 7, 8]
val = 8
print(binary(arr, val))
