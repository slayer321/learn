def solve(s):
    if (s == None or len(s) < 1):
        return ""
    res = ""
    for i in range(len(s)):
        len1 = helper(s, i, i)
        if len(len1) > len(res):
            res = len1
        len2 = helper(s, i, i+1)
        if len(len2) > len(res):
            res = len2
        #res = max(len1, len2)

    return res


def helper(s, l, r):
    if (s == None or l > r):
        return 0

    while (l >= 0 and r < len(s) and s[l] == s[r]):
        l -= 1
        r += 1

    return s[l+1:r]


s = "aaaabbaa"
print(solve(s))
