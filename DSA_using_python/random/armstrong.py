
def order(x):
    n = 0
    while x:
        n += 1
        x = x // 10
    return n


def solve(n):
    count = order(n)
    su = 0
    temp = n
    while temp:
        num = temp % 10
        su += pow(num, count)
        temp = temp // 10

    if n == su:
        res = "yes"
    else:
        res = "no"

    return res


if __name__ == "__main__":
    n = 153
    print(solve(n))
