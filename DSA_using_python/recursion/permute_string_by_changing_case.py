# https://www.geeksforgeeks.org/permute-string-changing-case/

def solve(s):
    op = ""
    ip = s
    res = []
    s1(ip, op, res)
    return res


def s1(ip, op, res):
    if len(ip) == 0:
        res.append(op)
        return

    op1 = op
    op2 = op
    op1 += ip[0]
    op2 += ip[0].upper()
    ip = ip[1:]
    s1(ip, op1, res)
    s1(ip, op2, res)


s = "a1b2"
print(solve(s))
