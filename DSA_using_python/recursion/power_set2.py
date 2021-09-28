def power_set(arr):
    ip = arr
    op = []
    res = []
    solve(ip, op, res)
    res.sort()
    # print(res)
    # print(len(res))
    out = []
    for i in range(len(res)-1):
        # print(res[i])
        if res[i] != res[i+1]:
            out.append(res[i+1])
    return out


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


s = [1, 2, 2]
print(power_set(s))
