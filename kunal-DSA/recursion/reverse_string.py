def solve(s):
    s1(s, 0, len(s)-1)


def s1(s, i, j):
    if i > j:
        return s
    s[i], s[j] = s[j], s[i]
    i += 1
    j -= 1
    return s1(s, i, j)


s = ["h", "e", "l", "l", "o"]
print(solve(s))
