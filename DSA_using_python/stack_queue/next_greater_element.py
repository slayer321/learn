def solve(arr, n):
    for i in range(n):
        for j in range(i, n):
            if arr[i] < arr[j]:
                arr[i] = arr[j]
                break
            elif j == n-1:
                arr[i] = -1

    return arr


def solve(arr,n):
    


arr = [6, 8, 0, 1, 3]
print(solve(arr, len(arr)))
