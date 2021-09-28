# def solve(arr):
#     n = len(arr)
#     for i in range(n):
#         for j in range(0, n-i-1):
#             if arr[j] > arr[j+1]:
#                 arr[j], arr[j+1] = arr[j+1], arr[j]
#     return arr


# recursive

def solve(arr, n):

    if n < 1:
        return arr
    for i in range(n-1):
        if arr[i] > arr[i+1]:
            arr[i], arr[i+1] = arr[i+1], arr[i]

    return solve(arr, n-1)


arr = [5, 1, 4, 2, 8]
print(solve(arr, len(arr)))
