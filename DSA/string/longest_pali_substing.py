'''
Longest Palindromic Substring
'''


def longestPalindrome(s):
    if len(s) <= 1:
        return s
    n = len(s)
    st, en = 0, 0
    max_len = 1

    # odd string
    if n % 2 != 0:
        for i in range(n-1):
            l = i
            r = i
            st, en = find_string(l, r, s, n, max_len, st, en)

    # Even string
    else:
        for i in range(n-1):
            l = i
            r = i + 1
            st, en = find_string(l, r, s, n, max_len, st, en)

    return s[st:en+1]


def find_string(l, r, s, n, max_len, st, en):
    while (l >= 0 and r < n):
        if s[l] == s[r]:
            l -= 1
            r += 1
        else:
            break

    val = r-l-1
    if val > max_len:
        max_len = val
        st = l + 1
        en = r - 1

    return st, en


s = "wtwosileez"
print(longestPalindrome(s))
