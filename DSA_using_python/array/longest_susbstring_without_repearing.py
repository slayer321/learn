# Longest Substring With Without Repeating Characters


def solve(s):
    i = 0
    j = 0
    dic = {}
    n = len(s)
    max_val = 0
    while j < n:
        if s[j] in dic:
            dic[s[j]] += 1
        else:
            dic[s[j]] = 1

        if len(dic) > j - i + 1:
            j += 1

        if len(dic) == j-i + 1:
            max_val = max(max_val, j-i+1)
            j += 1
        elif len(dic) < j-i+1:
            while len(dic) < j-i+1:
                dic[s[i]] -= 1
                if dic[s[i]] == 0:
                    del dic[s[i]]
                i += 1
            j += 1

    return max_val


s = "abcabcbb"
print(solve(s))
