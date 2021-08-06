'''
Delete the mid in stack

'''


def deleteMid(s, n):
    if n == 0:
        return s
    k = (n // 2) + 1
    solve(s, k)
    return s


def solve(s, k):
    if k == 1:
        s.pop()
        return s
    val = s[-1]
    s.pop()
    solve(s, k - 1)
    return s.append(val)


s = [1, 2, 3, 4]
print(deleteMid(s, len(s)))
