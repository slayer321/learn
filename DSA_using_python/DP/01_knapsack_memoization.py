#arr = [[-1, -1, -1, -1]] * 5
wt = [1, 3, 4, 5]
val = [1, 4, 5, 7]
W = 7
n = len(wt)

arr = [[-1] * 5] * 8
# for i in range(4):
#     for j in range(4):
#         arr[i][j] = 0

# print(arr)


def knapsack(wt, val, w, n):
    if (w == 0 or n == 0):
        return 0
    if arr[n][w] != -1:
        return arr[n][w]
    if wt[n-1] <= w:
        arr[n][w] = max(val[n-1] + knapsack(wt, val, w-wt[n-1],
                        n-1), knapsack(wt, val, w, n-1))
        return

    elif(wt[n-1] > w):
        arr[n][w] = knapsack(wt, val, w, n-1)
        return


print(knapsack(wt, val, W, n))
