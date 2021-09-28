def power_set(s):
    ip = s
    op = ""
    solve(ip, op)
    return


def solve(ip, op):
    if len(ip) == 0:
        print(op, end=",")
        return

    op1 = op
    op2 = op

    op2 += ip[0]
    ip = ip[1:]

    solve(ip, op1)
    solve(ip, op2)


s = "ab"
power_set(s)
