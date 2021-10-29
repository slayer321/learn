
def solve(wt, val, w, n):
    if (n == 0 or w == 0):
        return 0

    if wt[n-1] <= w:
        return max(val[n-1]+solve(wt, val, w-wt[n-1], n-1), solve(wt, val, w, n-1))
    elif wt[n-1] > w:
        return solve(wt, val, w, n-1)


wt = [1, 3, 4, 5]
val = [1, 4, 5, 7]
W = 7

print(solve(wt, val, W, len(wt)))
