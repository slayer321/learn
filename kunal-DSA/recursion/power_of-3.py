def solve(n):
    if n == 1:
        return True
    if n % 3 != 0 or n == 0:
        return False
    return solve(n/3)


n = 21
print(solve(n))
