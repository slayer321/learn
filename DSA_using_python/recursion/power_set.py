def power_set(arr):
    ip = arr
    op = []
    res = []
    solve(ip, op, res)
    return res


def solve(ip, op, res):
    if len(ip) == 0:
        res.append(op)
        return

    op1 = op[:]
    op2 = op[:]

    op2.append(ip[0])
    ip = ip[1:]

    solve(ip, op1, res)
    solve(ip, op2, res)


s = [1, 2, 3]
print(power_set(s))
