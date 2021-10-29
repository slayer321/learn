# https: // leetcode.com/problems/letter-case-permutation/


def solve(s):
    ip = s
    op = ""
    res = []
    s1(ip, op, res)

    return res


def s1(ip, op, res):
    if len(ip) == 0:
        res.append(op)
        return

    if ip[0].isalpha():
        op1 = op
        op2 = op
        op1 += ip[0].lower()

        op2 += ip[0].upper()
        ip = ip[1:]
        s1(ip, op1, res)
        s1(ip, op2, res)

    else:
        op1 = op
        op1 += ip[0]
        ip = ip[1:]
        s1(ip, op1, res)


s = "a1b2"
print(solve(s))
