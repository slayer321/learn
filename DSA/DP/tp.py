def tp(n):
    if n <= 0:
        return 0
    return max(n+tp(n-1), tp(n-2))


print(tp(5))
