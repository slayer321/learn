def f1(m, n):
    res = '0'
    out = f2(m, n, res)
    return out


def f2(m, n, res):
    m = m * 10
    n = n * 10
    res = str(m) + str(n)
    return res


m = 1
n = 1

print(f1(m, n))
