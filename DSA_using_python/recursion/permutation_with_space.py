def solve(s):
    ip = s
    op = ""
    op += op+s[0]
    ip = ip[1:]
    res = []
    sl(ip, op, res)
    return res


def sl(ip, op, res):
    if len(ip) == 0:
        #val = tuple(op)
        res.append(op)
        return
    op1 = op
    op2 = op
    op1 += ip[0]
    op2 += " " + ip[0]
    ip = ip[1:]
    sl(ip, op1, res)
    sl(ip, op2, res)


s = "ABC"
print(solve(s))
