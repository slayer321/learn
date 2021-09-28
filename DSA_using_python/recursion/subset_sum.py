def solve(arr, n):
    ip = arr
    op = 0
    res = []
    s(ip, op, res)
    return res


def s(ip, op, res):
    if len(ip) == 0:
        res.append(op)
        return
    op1 = op
    op2 = op

    op2 = op + ip[0]
    ip = ip[1:]
    s(ip, op1, res)
    s(ip, op2, res)


arr = [5, 2, 1]
print(solve(arr, len(arr)))
