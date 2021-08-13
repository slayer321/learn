
# Longest K unique characters substring


def longestKSubstr(s, k):
    # code here
    i = 0
    j = 0
    max_val = -1
    dic = {}
    n = len(s)
    while j < n:
        if s[j] in dic:
            dic[s[j]] += 1
        else:
            dic[s[j]] = 1

        if len(dic) < k:
            j += 1

        if len(dic) == k:
            max_val = max(max_val, j-i+1)
            j += 1
        elif len(dic) > k:
            while (len(dic) > k):
                dic[s[i]] -= 1
                if dic[s[i]] == 0:
                    del dic[s[i]]
                i += 1
            j += 1

    return max_val


s = "aabacbebebe"
k = 3
print(longestKSubstr(s, k))
