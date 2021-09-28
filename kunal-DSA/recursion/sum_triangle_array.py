def solve(arr):
    if len(arr) == 1:
        return
    for i in range(len(arr)-1):
        arr[i] += arr[i+1]
    arr.pop()
    solve(arr)
    print(arr)


arr = [1, 2, 3, 4, 5]
solve(arr)
