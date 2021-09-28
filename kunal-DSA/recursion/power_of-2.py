def pow2(n):
    return solve(n, 0)


def solve(n, i):
    if n == pow(2, i):
        return True
    if n == i:
        return False
    solve(n, i+1)


n = 16
print(pow2(n))
