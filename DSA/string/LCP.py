'''
Longest Common Prefix
'''


def LCP(s):
    if len(s) == 0:
        return s
    s1, s2 = min(s), max(s)
    i = 0
    while (i < len(s1) and i < len(s2) and s1[i] == s2[i]):
        i += 1
    return s1[:i]


s = ["flower", "flow", "flight"]
print(LCP(s))
